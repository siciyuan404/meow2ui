import React from 'react'
import type { AgentMessage, AgentMonitor, AgentRunOutput } from './agent-types'

type SaveTemplateStatus = 'idle' | 'saving' | 'success' | 'error'

export function useAgentRuntime() {
  const [sessionId, setSessionId] = React.useState('')
  const [inputText, setInputText] = React.useState('')
  const [messages, setMessages] = React.useState<AgentMessage[]>([])
  const [runStatus, setRunStatus] = React.useState<'idle' | 'running' | 'success' | 'error'>('idle')
  const [currentRunId, setCurrentRunId] = React.useState('')
  const [lastPrompt, setLastPrompt] = React.useState('')
  const [errorText, setErrorText] = React.useState('')
  const [traceId, setTraceId] = React.useState('')
  const [lastRunOutput, setLastRunOutput] = React.useState<AgentRunOutput | null>(null)
  const [saveTemplateStatus, setSaveTemplateStatus] = React.useState<SaveTemplateStatus>('idle')
  const [saveTemplateMessage, setSaveTemplateMessage] = React.useState('')
  const [monitor, setMonitor] = React.useState<AgentMonitor>({
    inputTokens: 0,
    outputTokens: 0,
    totalCost: 0,
    thinking: [],
    files: [],
    outputPreview: ['{', '  "type": "frame",', '  "layout": "vertical",'],
  })

  React.useEffect(() => {
    void (async () => {
      try {
        const wsRes = await fetch('/workspace/create', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ name: 'web-agent', root: '/tmp/web-agent' }),
        })
        const workspace = (await wsRes.json()) as { id?: string }
        if (!workspace.id) {
          return
        }
        const ssnRes = await fetch('/session/create', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ workspaceId: workspace.id, title: 'agent-session', themeId: 'default' }),
        })
        const sessionData = (await ssnRes.json()) as { session?: { id?: string } }
        if (sessionData.session?.id) {
          setSessionId(sessionData.session.id)
        }
      } catch {
        setErrorText('初始化会话失败，请刷新重试。')
      }
    })()
  }, [])

  const refreshMonitor = React.useCallback(async (runId: string, fallbackSchemaJSON?: string) => {
    try {
      const [detailRes, costRes] = await Promise.all([
        fetch(`/api/v1/debug/runs/${runId}`),
        fetch(`/api/v1/debug/runs/${runId}/cost`),
      ])
      const detail = (await detailRes.json()) as {
        traceId?: string
        steps?: Array<{ step?: string; tokenIn?: number; tokenOut?: number; payload?: Record<string, unknown> }>
      }
      const cost = (await costRes.json()) as { totalTokens?: number; totalCost?: number }

      const steps = detail.steps ?? []
      const thinking = steps.map((s) => `${s.step ?? 'step'}...`).slice(0, 4)
      const files = steps
        .map((s) => {
          const payload = (s.payload ?? {}) as Record<string, unknown>
          const v = payload.version_id
          return typeof v === 'string' ? `version-${v}.json` : ''
        })
        .filter(Boolean)
      const tokenIn = steps.reduce((acc, s) => acc + (s.tokenIn ?? 0), 0)
      const tokenOut = steps.reduce((acc, s) => acc + (s.tokenOut ?? 0), 0)

      setTraceId(detail.traceId ?? '')
      setMonitor({
        inputTokens: tokenIn,
        outputTokens: tokenOut,
        totalCost: cost.totalCost ?? 0,
        thinking,
        files: files.length > 0 ? files : ['dashboard.json'],
        outputPreview: fallbackSchemaJSON
          ? fallbackSchemaJSON.split('\n').slice(0, 3)
          : ['{', '  "type": "frame",', '  "layout": "vertical",'],
      })
    } catch {
      setMonitor((prev) => ({ ...prev, thinking: ['monitor data unavailable'] }))
    }
  }, [])

  const submitPrompt = React.useCallback(async (prompt: string) => {
    const text = prompt.trim()
    if (!text) {
      setErrorText('请输入 prompt 后再发送。')
      return
    }
    if (!sessionId) {
      setErrorText('会话尚未初始化完成，请稍后重试。')
      return
    }

    const userMsg: AgentMessage = { id: `u-${Date.now()}`, role: 'user', text }
    const pendingMsg: AgentMessage = { id: `a-${Date.now()}`, role: 'assistant', text: '正在生成中...', status: 'running' }
    setMessages((prev) => [...prev, userMsg, pendingMsg])
    setInputText('')
    setErrorText('')
    setRunStatus('running')
    setLastPrompt(text)

    try {
      const res = await fetch('/agent/run', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ sessionId, prompt: text, onlyArea: 'main' }),
      })
      const data = (await res.json()) as AgentRunOutput & { error?: string; message?: string }
      if (!res.ok || !data.runId) {
        const msg = data.error || data.message || '运行失败'
        throw new Error(msg)
      }

      setCurrentRunId(data.runId)
      setLastRunOutput(data)
      setRunStatus('success')
      setMessages((prev) =>
        prev.map((m) =>
          m.id === pendingMsg.id
            ? {
                ...m,
                text: `运行完成 run=${data.runId} version=${data.versionId}${data.repaired ? ' (repaired)' : ''}`,
                status: 'success',
              }
            : m,
        ),
      )
      await refreshMonitor(data.runId, data.schemaJSON)
    } catch (err) {
      const msg = err instanceof Error ? err.message : '运行失败'
      setRunStatus('error')
      setErrorText(msg)
      setMessages((prev) =>
        prev.map((m) => (m.id === pendingMsg.id ? { ...m, text: `运行失败: ${msg}`, status: 'error' } : m)),
      )
    }
  }, [refreshMonitor, sessionId])

  const rerunLast = React.useCallback(() => {
    if (lastPrompt && runStatus !== 'running') {
      void submitPrompt(lastPrompt)
    }
  }, [lastPrompt, runStatus, submitPrompt])

  const saveLastResultAsTemplate = React.useCallback(async () => {
    if (!lastRunOutput) {
      setSaveTemplateStatus('error')
      setSaveTemplateMessage('暂无可保存结果，请先运行一次。')
      return
    }
    setSaveTemplateStatus('saving')
    setSaveTemplateMessage('')
    try {
      const res = await fetch('/api/v1/marketplace/templates', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          name: `Agent Result ${lastRunOutput.runId.slice(-6)}`,
          category: 'dashboard',
          tags: ['agent', 'generated'],
          schema: lastRunOutput.schemaJSON,
          theme: 'default',
          owner: 'web-user',
          sessionId,
          versionId: lastRunOutput.versionId,
        }),
      })
      const data = (await res.json()) as { id?: string; error?: string; message?: string }
      if (!res.ok || !data.id) {
        const msg = data.error || data.message || '保存失败'
        throw new Error(msg)
      }
      setSaveTemplateStatus('success')
      setSaveTemplateMessage(`模板已保存 (${data.id})，可前往 Playground 查看。`)
    } catch (err) {
      setSaveTemplateStatus('error')
      setSaveTemplateMessage(err instanceof Error ? err.message : '保存失败')
    }
  }, [lastRunOutput, sessionId])

  const onEditorKeyDown: React.KeyboardEventHandler<HTMLTextAreaElement> = (e) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      if (runStatus !== 'running') {
        void submitPrompt(inputText)
      }
    }
  }

  return {
    sessionId,
    inputText,
    setInputText,
    messages,
    runStatus,
    currentRunId,
    errorText,
    traceId,
    monitor,
    lastRunOutput,
    saveTemplateStatus,
    saveTemplateMessage,
    submitPrompt,
    rerunLast,
    saveLastResultAsTemplate,
    onEditorKeyDown,
  }
}
