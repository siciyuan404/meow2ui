import React from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Link, Route, Routes, useLocation } from 'react-router-dom'
import { useAgentRuntime } from './use-agent-runtime'

const FONT_UI = 'IBM Plex Mono, monospace'
const FONT_DISPLAY = 'JetBrains Mono, monospace'

type IconName = 'workspaces' | 'playground' | 'agent' | 'providerPool' | 'debug' | 'settings' | 'search' | 'close'

function AppIcon({ name, size = 18 }: { name: IconName; size?: number }) {
  const common = {
    width: size,
    height: size,
    viewBox: '0 0 24 24',
    fill: 'none',
    stroke: 'currentColor',
    strokeWidth: 1.8,
    strokeLinecap: 'round' as const,
    strokeLinejoin: 'round' as const,
    'aria-hidden': true,
  }

  switch (name) {
    case 'workspaces':
      return (
        <svg {...common}>
          <rect x="3" y="4" width="8" height="7" rx="1.5" />
          <rect x="13" y="4" width="8" height="7" rx="1.5" />
          <rect x="3" y="13" width="8" height="7" rx="1.5" />
          <rect x="13" y="13" width="8" height="7" rx="1.5" />
        </svg>
      )
    case 'playground':
      return (
        <svg {...common}>
          <path d="M8 3v6" />
          <path d="M16 3v6" />
          <path d="M3 10h18" />
          <rect x="3" y="6" width="18" height="15" rx="2" />
        </svg>
      )
    case 'providerPool':
      return (
        <svg {...common}>
          <circle cx="8" cy="12" r="3" />
          <path d="M11 12h10" />
          <path d="M18 9v6" />
          <path d="M5 12H3" />
        </svg>
      )
    case 'agent':
      return (
        <svg {...common}>
          <rect x="5" y="7" width="14" height="11" rx="2" />
          <circle cx="9" cy="12" r="1.5" />
          <circle cx="15" cy="12" r="1.5" />
          <path d="M12 4v3" />
          <path d="M9 18h6" />
        </svg>
      )
    case 'debug':
      return (
        <svg {...common}>
          <path d="M12 5v14" />
          <path d="M5 12h14" />
          <circle cx="12" cy="12" r="8" />
        </svg>
      )
    case 'settings':
      return (
        <svg {...common}>
          <circle cx="12" cy="12" r="3" />
          <path d="M19.4 15a1 1 0 0 0 .2 1.1l.1.1a1 1 0 0 1-1.4 1.4l-.1-.1a1 1 0 0 0-1.1-.2 1 1 0 0 0-.6.9V19a1 1 0 1 1-2 0v-.2a1 1 0 0 0-.6-.9 1 1 0 0 0-1.1.2l-.1.1a1 1 0 0 1-1.4-1.4l.1-.1a1 1 0 0 0 .2-1.1 1 1 0 0 0-.9-.6H5a1 1 0 1 1 0-2h.2a1 1 0 0 0 .9-.6 1 1 0 0 0-.2-1.1l-.1-.1a1 1 0 0 1 1.4-1.4l.1.1a1 1 0 0 0 1.1.2 1 1 0 0 0 .6-.9V5a1 1 0 1 1 2 0v.2a1 1 0 0 0 .6.9 1 1 0 0 0 1.1-.2l.1-.1a1 1 0 0 1 1.4 1.4l-.1.1a1 1 0 0 0-.2 1.1 1 1 0 0 0 .9.6H19a1 1 0 1 1 0 2h-.2a1 1 0 0 0-.9.6Z" />
        </svg>
      )
    case 'search':
      return (
        <svg {...common}>
          <circle cx="11" cy="11" r="6" />
          <path d="m20 20-4.2-4.2" />
        </svg>
      )
    case 'close':
      return (
        <svg {...common}>
          <path d="M18 6 6 18" />
          <path d="m6 6 12 12" />
        </svg>
      )
  }
}

function Page({ title }: { title: string }) {
  const routeCta: Record<string, { title: string; desc: string; actionLabel: string; actionPath: string }> = {
    Workspaces: {
      title: '还没有工作区',
      desc: '创建一个工作区，开始你的第一个 UI 生成任务。',
      actionLabel: '前往 Agent',
      actionPath: '/agent',
    },
    Playground: {
      title: '还没有模板',
      desc: '先在 Agent 生成一次，再一键保存到模板市场。',
      actionLabel: '前往 Agent',
      actionPath: '/agent',
    },
    Settings: {
      title: '还没有配置项',
      desc: '可先去凭证池配置 Provider，提升生成质量。',
      actionLabel: '打开凭证池',
      actionPath: '/provider-pool',
    },
  }
  const empty = routeCta[title]
  return (
    <div style={{ padding: 24 }}>
      <h2 style={{ fontFamily: FONT_DISPLAY, marginTop: 0 }}>{title}</h2>
      {empty ? (
        <div style={{ border: '1px solid #e5e7eb', borderRadius: 12, background: '#fff', padding: 16, maxWidth: 520 }}>
          <div style={{ fontFamily: FONT_DISPLAY, fontSize: 16, marginBottom: 6 }}>{empty.title}</div>
          <div style={{ color: '#64748b', marginBottom: 12 }}>{empty.desc}</div>
          <Link to={empty.actionPath} style={{ display: 'inline-block', padding: '8px 12px', borderRadius: 8, background: '#10b981', color: '#0a0a0a', fontWeight: 700 }}>
            {empty.actionLabel}
          </Link>
        </div>
      ) : (
        <p>Coming soon</p>
      )}
    </div>
  )
}

type DebugRun = {
  runId: string
  sessionId: string
  status: string
  startedAt: string
}

type DebugDetail = {
  run: DebugRun
  steps: Array<{ step: string; status: string; latencyMs: number; tokenIn: number; tokenOut: number }>
  cost: { totalTokens: number; totalCost: number; byProvider: Record<string, number>; byModel: Record<string, number> }
}

function DebugPage() {
  const [runs, setRuns] = React.useState<DebugRun[]>([])
  const [selected, setSelected] = React.useState<string>('')
  const [detail, setDetail] = React.useState<DebugDetail | null>(null)
  const [loading, setLoading] = React.useState(false)

  React.useEffect(() => {
    void fetch('/api/v1/debug/runs')
      .then((r) => r.json())
      .then((data) => {
        const items = (data?.runs ?? []) as DebugRun[]
        setRuns(items)
        if (items.length > 0) {
          setSelected(items[0].runId)
        }
      })
      .catch(() => {
        setRuns([])
      })
  }, [])

  React.useEffect(() => {
    if (!selected) {
      setDetail(null)
      return
    }
    setLoading(true)
    void fetch(`/api/v1/debug/runs/${selected}`)
      .then((r) => r.json())
      .then((data) => {
        setDetail(data as DebugDetail)
      })
      .catch(() => {
        setDetail(null)
      })
      .finally(() => {
        setLoading(false)
      })
  }, [selected])

  return (
    <div style={{ display: 'grid', gridTemplateColumns: '320px 1fr', gap: 16, padding: 16 }}>
      <div style={{ gridColumn: '1 / -1', display: 'grid', gridTemplateColumns: 'repeat(3,minmax(0,1fr))', gap: 12 }}>
        <div style={{ border: '1px solid #e5e7eb', borderRadius: 8, background: '#fff', padding: 12 }}>
          <div style={{ color: '#64748b', fontSize: 12 }}>Success Rate</div>
          <div style={{ fontFamily: FONT_DISPLAY, fontSize: 20, fontWeight: 700 }}>
            {runs.length > 0 ? `${Math.round((runs.filter((r) => r.status === 'completed').length / runs.length) * 100)}%` : '0%'}
          </div>
        </div>
        <div style={{ border: '1px solid #e5e7eb', borderRadius: 8, background: '#fff', padding: 12 }}>
          <div style={{ color: '#64748b', fontSize: 12 }}>Avg Latency</div>
          <div style={{ fontFamily: FONT_DISPLAY, fontSize: 20, fontWeight: 700 }}>
            {detail ? `${Math.max(0, Math.round(detail.steps.reduce((a, s) => a + s.latencyMs, 0) / Math.max(1, detail.steps.length)))}ms` : '--'}
          </div>
        </div>
        <div style={{ border: '1px solid #e5e7eb', borderRadius: 8, background: '#fff', padding: 12 }}>
          <div style={{ color: '#64748b', fontSize: 12 }}>Total Cost</div>
          <div style={{ fontFamily: FONT_DISPLAY, fontSize: 20, fontWeight: 700 }}>
            {detail ? `$${detail.cost.totalCost.toFixed(6)}` : '$0.000000'}
          </div>
        </div>
      </div>
      <div style={{ border: '1px solid #ddd', borderRadius: 8, padding: 12 }}>
        <h3 style={{ marginTop: 0 }}>Agent Runs</h3>
        {runs.length === 0 ? <p>No runs</p> : null}
        {runs.map((item) => (
          <button
            key={item.runId}
            type="button"
            onClick={() => setSelected(item.runId)}
            style={{
              width: '100%',
              textAlign: 'left',
              marginBottom: 8,
              padding: 8,
              border: selected === item.runId ? '1px solid #333' : '1px solid #ddd',
              borderRadius: 6,
              background: '#fff',
              cursor: 'pointer',
            }}
          >
            <div style={{ fontFamily: FONT_DISPLAY }}><strong>{item.runId}</strong></div>
            <div>session: {item.sessionId}</div>
            <div>status: {item.status}</div>
          </button>
        ))}
      </div>

      <div style={{ border: '1px solid #ddd', borderRadius: 8, padding: 12 }}>
        <h3 style={{ marginTop: 0, fontFamily: FONT_DISPLAY }}>Run Detail</h3>
        {loading ? <p>Loading...</p> : null}
        {!loading && !detail ? <p>Select a run.</p> : null}
        {detail ? (
          <>
            <div style={{ marginBottom: 12 }}>
              <div><strong>Run:</strong> {detail.run.runId}</div>
              <div><strong>Status:</strong> {detail.run.status}</div>
              <div><strong>Started:</strong> {detail.run.startedAt}</div>
            </div>
            <h4 style={{ fontFamily: FONT_DISPLAY }}>Steps</h4>
            <table style={{ width: '100%', borderCollapse: 'collapse', marginBottom: 12 }}>
              <thead>
                <tr>
                  <th style={{ textAlign: 'left' }}>step</th>
                  <th style={{ textAlign: 'left' }}>status</th>
                  <th style={{ textAlign: 'left' }}>latency</th>
                  <th style={{ textAlign: 'left' }}>tokens(in/out)</th>
                </tr>
              </thead>
              <tbody>
                {detail.steps.map((s, i) => (
                  <tr key={`${s.step}-${i}`}>
                    <td>{s.step}</td>
                    <td style={{ color: s.status === 'failed' ? '#b42318' : '#067647' }}>{s.status}</td>
                    <td>{s.latencyMs}ms</td>
                    <td>{s.tokenIn}/{s.tokenOut}</td>
                  </tr>
                ))}
              </tbody>
            </table>
            <h4 style={{ fontFamily: FONT_DISPLAY }}>Cost</h4>
            <div>tokens: {detail.cost.totalTokens}</div>
            <div>total: ${detail.cost.totalCost.toFixed(6)}</div>
          </>
        ) : null}
      </div>
    </div>
  )
}

function IconSidebar() {
  const { pathname } = useLocation()
  const iconLink = (icon: IconName, to: string) => {
    const active = pathname === to
    return (
      <Link
        to={to}
        aria-label={to}
        title={to}
        style={{
          width: 38,
          height: 38,
          borderRadius: 10,
          textDecoration: 'none',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          color: active ? '#0a0a0a' : '#6b7280',
          background: active ? '#10b981' : 'transparent',
          fontSize: 13,
          fontWeight: 700,
          fontFamily: FONT_DISPLAY,
        }}
      >
        <AppIcon name={icon} size={18} />
      </Link>
    )
  }

  return (
    <aside
      style={{
        width: 54,
        minWidth: 54,
        background: '#0f0f0f',
        borderRight: '1px solid #1f1f1f',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        padding: '12px 0',
        gap: 8,
      }}
    >
      <div
        style={{
          width: 40,
          height: 40,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        <div
          style={{
            width: 34,
            height: 34,
            borderRadius: 6,
            background: '#10b981',
            color: '#0a0a0a',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            fontFamily: FONT_DISPLAY,
            fontSize: 25,
            fontWeight: 700,
            lineHeight: 1,
          }}
        >
          M
        </div>
      </div>

      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          gap: 4,
          flex: 1,
          width: '100%',
          padding: '16px 0',
        }}
      >
        {iconLink('workspaces', '/workspaces')}
        {iconLink('playground', '/playground')}
        {iconLink('agent', '/agent')}
        {iconLink('providerPool', '/provider-pool')}
      </div>

      <div
        style={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          gap: 4,
          width: '100%',
          paddingTop: 8,
          borderTop: '1px solid #1f1f1f',
        }}
      >
        {iconLink('debug', '/debug')}
        {iconLink('settings', '/settings')}
      </div>
    </aside>
  )
}

function ProviderPoolPage() {
  const [activeProvider, setActiveProvider] = React.useState('OpenAI')
  const [testStatus, setTestStatus] = React.useState<'idle' | 'loading' | 'success' | 'error'>('idle')
  const [testMessage, setTestMessage] = React.useState('')

  const runConnectionTest = React.useCallback(async () => {
    setTestStatus('loading')
    setTestMessage('连接测试中...')
    await new Promise((resolve) => setTimeout(resolve, 400))
    if (activeProvider === 'OpenAI') {
      setTestStatus('success')
      setTestMessage('连接成功，可正常调用。')
    } else {
      setTestStatus('error')
      setTestMessage('连接失败，请检查 key 或 base url。')
    }
  }, [activeProvider])

  return (
    <div style={{ height: 'calc(100vh - 0px)', background: '#0f0f0f', color: '#fafafa' }}>
      <div
        style={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
          padding: '16px 24px',
          borderBottom: '1px solid #1f1f1f',
        }}
      >
        <div style={{ display: 'flex', alignItems: 'center', gap: 12 }}>
          <span style={{ color: '#10b981', display: 'inline-flex' }}><AppIcon name="providerPool" size={20} /></span>
          <span style={{ fontFamily: FONT_DISPLAY, fontSize: 16, fontWeight: 600 }}>凭证池 - API Key 管理</span>
          <span style={{ color: '#4b5563', fontFamily: FONT_UI, fontSize: 11 }}>/provider-pool</span>
        </div>
        <span style={{ color: '#6b7280', display: 'inline-flex' }}><AppIcon name="close" size={16} /></span>
      </div>

      <div style={{ display: 'flex', height: 'calc(100vh - 57px)' }}>
        <div style={{ width: 260, borderRight: '1px solid #1f1f1f', display: 'flex', flexDirection: 'column' }}>
          <div style={{ padding: '12px 16px', borderBottom: '1px solid #1f1f1f' }}>
            <div
              style={{
                display: 'flex',
                alignItems: 'center',
                gap: 8,
                borderRadius: 8,
                border: '1px solid #1f1f1f',
                background: '#111827',
                color: '#9ca3af',
                fontFamily: FONT_UI,
                fontSize: 12,
                padding: '10px 12px',
              }}
            >
              <AppIcon name="search" size={14} />
              搜索 Provider / Key
            </div>
          </div>
          <div style={{ flex: 1, overflow: 'auto', padding: 8 }}>
            <div style={{ color: '#9ca3af', fontSize: 12, fontFamily: FONT_UI, padding: '8px 10px' }}>OpenAI Compatible</div>
            <button type="button" onClick={() => setActiveProvider('DeepSeek')} style={{ width: '100%', textAlign: 'left', border: 'none', borderRadius: 8, background: activeProvider === 'DeepSeek' ? '#10b981' : '#111827', color: activeProvider === 'DeepSeek' ? '#0a0a0a' : '#fafafa', padding: '10px 12px', marginBottom: 6, fontWeight: activeProvider === 'DeepSeek' ? 700 : 400 }}>DeepSeek</button>
            <button type="button" onClick={() => setActiveProvider('OpenAI')} style={{ width: '100%', textAlign: 'left', border: 'none', borderRadius: 8, background: activeProvider === 'OpenAI' ? '#10b981' : '#111827', color: activeProvider === 'OpenAI' ? '#0a0a0a' : '#fafafa', padding: '10px 12px', marginBottom: 6, fontWeight: activeProvider === 'OpenAI' ? 700 : 400 }}>OpenAI</button>
            <button type="button" onClick={() => setActiveProvider('Anthropic')} style={{ width: '100%', textAlign: 'left', border: 'none', borderRadius: 8, background: activeProvider === 'Anthropic' ? '#10b981' : '#111827', color: activeProvider === 'Anthropic' ? '#0a0a0a' : '#fafafa', padding: '10px 12px', marginBottom: 6, fontWeight: activeProvider === 'Anthropic' ? 700 : 400 }}>Anthropic</button>
          </div>
          <div style={{ borderTop: '1px solid #1f1f1f', padding: 12, display: 'flex', gap: 8 }}>
            <button type="button" style={{ flex: 1, background: '#10b981', color: '#0a0a0a', border: 'none', borderRadius: 8, padding: '8px 10px', fontWeight: 700 }}>新增</button>
            <button type="button" style={{ flex: 1, background: '#1f2937', color: '#d1d5db', border: '1px solid #374151', borderRadius: 8, padding: '8px 10px' }}>删除</button>
          </div>
        </div>

        <div style={{ flex: 1, padding: 24, overflow: 'auto' }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 12 }}>
            <h3 style={{ margin: 0, fontFamily: FONT_DISPLAY }}>Provider Settings</h3>
            <span style={{ color: '#9ca3af', fontFamily: FONT_UI, fontSize: 12 }}>{activeProvider} (active)</span>
          </div>
          <div style={{ height: 1, background: '#1f1f1f', marginBottom: 20 }} />

          <div style={{ display: 'grid', gap: 12, marginBottom: 20 }}>
            <label style={{ color: '#9ca3af', fontSize: 12 }}>API Key</label>
            <input
              placeholder="sk-..."
              style={{ background: '#111827', color: '#fafafa', border: '1px solid #374151', borderRadius: 8, padding: '10px 12px' }}
            />
            <label style={{ color: '#9ca3af', fontSize: 12 }}>Base URL</label>
            <input
              placeholder="https://api.openai.com/v1"
              style={{ background: '#111827', color: '#fafafa', border: '1px solid #374151', borderRadius: 8, padding: '10px 12px' }}
            />
          </div>

          <div style={{ display: 'flex', gap: 8 }}>
            <button type="button" style={{ background: '#10b981', color: '#0a0a0a', border: 'none', borderRadius: 8, padding: '10px 14px', fontWeight: 700 }}>保存配置</button>
            <button type="button" onClick={() => void runConnectionTest()} style={{ background: '#1f2937', color: '#d1d5db', border: '1px solid #374151', borderRadius: 8, padding: '10px 14px' }}>
              {testStatus === 'loading' ? '测试中...' : '连接测试'}
            </button>
          </div>
          {testMessage ? (
            <div style={{ marginTop: 10, color: testStatus === 'success' ? '#34d399' : testStatus === 'error' ? '#fca5a5' : '#9ca3af', fontSize: 12 }}>
              {testMessage}
            </div>
          ) : null}
        </div>
      </div>
    </div>
  )
}

function AgentPage() {
  const {
    sessionId,
    inputText,
    setInputText,
    messages,
    runStatus,
    currentRunId,
    errorText,
    traceId,
    monitor,
    submitPrompt,
    rerunLast,
    saveLastResultAsTemplate,
    saveTemplateStatus,
    saveTemplateMessage,
    onEditorKeyDown,
  } = useAgentRuntime()

  const starterPrompts = [
    { label: 'Dashboard', prompt: '创建一个现代 SaaS 仪表盘，包含统计卡片、趋势图和最近活动列表。' },
    { label: 'Form', prompt: '创建一个多步骤表单页面，包含基本信息、联系方式和确认提交。' },
    { label: 'Landing', prompt: '创建一个产品落地页，包含 Hero、特性区、用户评价和 CTA。' },
  ]

  return (
    <div style={{ height: '100vh', display: 'flex', background: '#0a0a0a', color: '#fafafa' }}>
      <section
        style={{
          width: 260,
          borderRight: '1px solid #1f1f1f',
          padding: '24px 16px',
          display: 'flex',
          flexDirection: 'column',
          gap: 16,
          background: '#0f0f0f',
        }}
      >
        <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
          <span style={{ color: '#10b981', fontFamily: FONT_DISPLAY, fontWeight: 700, fontSize: 20 }}>{'>'}</span>
          <span style={{ fontFamily: FONT_DISPLAY, fontSize: 18, fontWeight: 600 }}>Meow2UI</span>
          <span style={{ width: 2, height: 20, background: '#10b981' }} />
        </div>
        <div style={{ height: 1, background: '#1f1f1f' }} />
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
          <span style={{ color: '#6b7280', fontSize: 11 }}>{'// workspaces'}</span>
          <span style={{ color: '#10b981', fontSize: 14 }}>+</span>
        </div>
        <div style={{ flex: 1, display: 'flex', flexDirection: 'column', gap: 6 }}>
          <div style={{ borderRadius: 8, background: '#1f1f1f' }}>
            <div style={{ display: 'flex', gap: 8, padding: '10px 12px', alignItems: 'center' }}>
              <span style={{ width: 6, height: 6, borderRadius: 6, background: '#10b981' }} />
              <span style={{ fontSize: 13 }}>my-saas-app/</span>
            </div>
            <div style={{ padding: '0 12px 10px 28px', display: 'flex', flexDirection: 'column', gap: 4 }}>
              <div style={{ borderRadius: 6, background: '#10b98120', padding: '6px 8px', fontSize: 12 }}>agent-dashboard</div>
              <div style={{ borderRadius: 6, padding: '6px 8px', color: '#9ca3af', fontSize: 12 }}>landing-v2</div>
            </div>
          </div>
          <div style={{ borderRadius: 6, padding: '10px 12px', color: '#6b7280', fontSize: 13 }}>e-commerce/</div>
          <div style={{ borderRadius: 6, padding: '10px 12px', color: '#6b7280', fontSize: 13 }}>portfolio/</div>
        </div>
        <div style={{ height: 1, background: '#1f1f1f' }} />
        <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
          <span style={{ color: '#6b7280', fontSize: 11 }}>{'// settings'}</span>
          <button type="button" style={{ textAlign: 'left', border: 'none', background: 'transparent', color: '#fafafa', padding: '8px 12px', borderRadius: 6 }}>AI Providers</button>
          <button type="button" style={{ textAlign: 'left', border: 'none', background: 'transparent', color: '#fafafa', padding: '8px 12px', borderRadius: 6 }}>Themes</button>
          <button type="button" style={{ textAlign: 'left', border: 'none', background: 'transparent', color: '#fafafa', padding: '8px 12px', borderRadius: 6 }}>Examples</button>
        </div>
      </section>

      <section style={{ flex: 1, display: 'flex', background: '#0f0f0f' }}>
        <div style={{ flex: 1, display: 'flex' }}>
          <div style={{ flex: 1, padding: 16 }}>
            <div style={{ height: '100%', border: '1px solid #1f1f1f', borderRadius: 8, background: '#0a0a0a', position: 'relative' }}>
              <div style={{ position: 'absolute', left: 16, top: 16, width: 200, height: 'calc(100% - 32px)', borderRadius: 8, border: '1px solid #2a2a2a', background: '#1a1a1a' }} />
              <div style={{ position: 'absolute', left: 225, top: 16, width: 40, height: 220, borderRadius: 8, border: '1px solid #2a2a2a', background: '#1a1a1a' }} />
              <div style={{ position: 'absolute', left: 280, top: 40, color: '#6b7280', fontSize: 12 }}>Canvas Preview</div>
            </div>
          </div>

          <div style={{ width: '32%', minWidth: 360, borderLeft: '1px solid #1f1f1f', display: 'flex', flexDirection: 'column', padding: 10, background: '#0a0a0a' }}>
            <div style={{ height: 52, borderBottom: '1px solid #1a1a1a', display: 'flex', alignItems: 'center', padding: '0 24px', fontFamily: FONT_DISPLAY, fontWeight: 600 }}>对话</div>
            <div style={{ flex: 1, padding: 20, display: 'flex', flexDirection: 'column', gap: 12, overflow: 'auto' }}>
              {messages.length === 0 ? (
                <div style={{ border: '1px solid #1f1f1f', borderRadius: 12, padding: 12, background: '#111111' }}>
                  <div style={{ color: '#e5e7eb', fontFamily: FONT_DISPLAY, marginBottom: 6 }}>从一个 Starter Prompt 开始</div>
                  <div style={{ color: '#6b7280', fontSize: 12, marginBottom: 10 }}>选择一个模板化需求，一键填充并发送。</div>
                  <div style={{ display: 'flex', gap: 8, flexWrap: 'wrap' }}>
                    {starterPrompts.map((s) => (
                      <button
                        key={s.label}
                        type="button"
                        onClick={() => {
                          setInputText(s.prompt)
                          void submitPrompt(s.prompt)
                        }}
                        disabled={runStatus === 'running' || !sessionId}
                        style={{
                          border: '1px solid #374151',
                          borderRadius: 8,
                          padding: '6px 10px',
                          background: '#0f172a',
                          color: '#d1d5db',
                          cursor: runStatus === 'running' || !sessionId ? 'not-allowed' : 'pointer',
                          opacity: runStatus === 'running' || !sessionId ? 0.6 : 1,
                        }}
                      >
                        {s.label}
                      </button>
                    ))}
                  </div>
                </div>
              ) : null}
              {messages.map((m) => {
                const isUser = m.role === 'user'
                const bg = isUser ? '#10b981' : m.status === 'error' ? '#7f1d1d' : '#1f1f1f'
                const color = isUser ? '#0a0a0a' : '#fafafa'
                return (
                  <div
                    key={m.id}
                    style={{
                      alignSelf: isUser ? 'flex-end' : 'flex-start',
                      borderRadius: 14,
                      padding: '8px 12px',
                      background: bg,
                      color,
                      maxWidth: '90%',
                      border: m.status === 'running' ? '1px dashed #10b981' : '1px solid transparent',
                    }}
                  >
                    {m.text}
                  </div>
                )
              })}
              {errorText ? <div style={{ color: '#fca5a5', fontSize: 12 }}>错误: {errorText}{traceId ? ` | trace: ${traceId}` : ''}</div> : null}
            </div>
            <div style={{ border: '1px solid #1f1f1f', borderRadius: 16, padding: 8, display: 'grid', gap: 8 }}>
              <textarea
                value={inputText}
                onChange={(e) => setInputText(e.target.value)}
                onKeyDown={onEditorKeyDown}
                placeholder="描述你想要生成或修改的 UI... (Enter 发送, Shift+Enter 换行)"
                style={{
                  minHeight: 56,
                  resize: 'vertical',
                  border: '1px solid #1f1f1f',
                  borderRadius: 10,
                  padding: '10px 12px',
                  color: '#e5e7eb',
                  background: '#0b0b0b',
                  outline: 'none',
                }}
              />
              <div style={{ display: 'flex', gap: 8, justifyContent: 'space-between' }}>
                <div style={{ display: 'flex', gap: 8 }}>
                  <button
                    type="button"
                    onClick={() => void submitPrompt(inputText)}
                    disabled={runStatus === 'running' || !sessionId}
                    style={{
                      background: '#10b981',
                      color: '#0a0a0a',
                      border: 'none',
                      borderRadius: 8,
                      padding: '8px 12px',
                      fontWeight: 700,
                      cursor: runStatus === 'running' ? 'not-allowed' : 'pointer',
                      opacity: runStatus === 'running' ? 0.6 : 1,
                    }}
                  >
                    {runStatus === 'running' ? '发送中...' : '发送'}
                  </button>
                  <button
                    type="button"
                    onClick={rerunLast}
                    disabled={runStatus === 'running'}
                    style={{
                      background: '#1f2937',
                      color: '#d1d5db',
                      border: '1px solid #374151',
                      borderRadius: 8,
                      padding: '8px 12px',
                      cursor: runStatus === 'running' ? 'not-allowed' : 'pointer',
                      opacity: runStatus === 'running' ? 0.6 : 1,
                    }}
                  >
                    重试上次
                  </button>
                </div>
                <button
                  type="button"
                  onClick={saveLastResultAsTemplate}
                  disabled={saveTemplateStatus === 'saving' || runStatus === 'running'}
                  style={{
                    background: saveTemplateStatus === 'success' ? '#065f46' : '#0f172a',
                    color: '#e2e8f0',
                    border: '1px solid #334155',
                    borderRadius: 8,
                    padding: '8px 12px',
                    cursor: saveTemplateStatus === 'saving' || runStatus === 'running' ? 'not-allowed' : 'pointer',
                    opacity: saveTemplateStatus === 'saving' || runStatus === 'running' ? 0.6 : 1,
                  }}
                >
                  {saveTemplateStatus === 'saving' ? '保存中...' : '保存模板'}
                </button>
              </div>
              {saveTemplateMessage ? (
                <div style={{ fontSize: 12, color: saveTemplateStatus === 'error' ? '#fca5a5' : '#86efac' }}>{saveTemplateMessage}</div>
              ) : null}
            </div>
          </div>
        </div>

        <aside style={{ width: 300, borderLeft: '1px solid #1f1f1f', padding: '20px 16px', display: 'flex', flexDirection: 'column', gap: 12 }}>
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <span style={{ color: '#6b7280', fontSize: 12 }}>{'// monitor'}</span>
            <span style={{ color: '#4b5563', fontSize: 11 }}>[x]</span>
          </div>
          <div style={{ height: 1, background: '#1f1f1f' }} />

          <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
            <span style={{ color: '#10b981', fontSize: 11 }}>$ tokens</span>
            <div style={{ display: 'flex', gap: 16 }}>
              <div><div style={{ color: '#4b5563', fontSize: 10 }}>input</div><div style={{ fontFamily: FONT_DISPLAY, fontSize: 16, fontWeight: 600 }}>{monitor.inputTokens}</div></div>
              <div><div style={{ color: '#4b5563', fontSize: 10 }}>output</div><div style={{ fontFamily: FONT_DISPLAY, fontSize: 16, fontWeight: 600 }}>{monitor.outputTokens}</div></div>
              <div><div style={{ color: '#4b5563', fontSize: 10 }}>cost</div><div style={{ fontFamily: FONT_DISPLAY, fontSize: 16, fontWeight: 600, color: '#10b981' }}>${monitor.totalCost.toFixed(4)}</div></div>
            </div>
          </div>

          <div style={{ height: 1, background: '#1f1f1f' }} />

          <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
            <span style={{ color: '#10b981', fontSize: 11 }}>$ thinking</span>
            <div style={{ border: '1px solid #1f1f1f', borderRadius: 8, background: '#0f0f0f', padding: 12, minHeight: 100, display: 'flex', flexDirection: 'column', gap: 4 }}>
              {(monitor.thinking.length > 0 ? monitor.thinking : ['等待运行数据...']).map((line, idx) => (
                <span key={`${line}-${idx}`} style={{ color: idx === (monitor.thinking.length - 1) ? '#fafafa' : '#6b7280', fontSize: 11 }}>{line}</span>
              ))}
            </div>
          </div>

          <div style={{ height: 1, background: '#1f1f1f' }} />

          <div style={{ display: 'flex', flexDirection: 'column', gap: 8, flex: 1 }}>
            <span style={{ color: '#10b981', fontSize: 11 }}>$ files</span>
            {monitor.files.map((f, idx) => (
              <div
                key={`${f}-${idx}`}
                style={{
                  borderRadius: 6,
                  background: idx === 0 ? '#10B98115' : idx === 1 ? '#F59E0B15' : 'transparent',
                  padding: '6px 8px',
                  fontSize: 11,
                  color: idx > 1 ? '#6b7280' : '#e5e7eb',
                }}
              >
                [{idx === 1 ? '~' : '+'}] {f}
              </div>
            ))}
          </div>

          <div style={{ height: 1, background: '#1f1f1f' }} />

          <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
            <span style={{ color: '#10b981', fontSize: 11 }}>$ output.json</span>
            <div style={{ border: '1px solid #1f1f1f', borderRadius: 8, background: '#0f0f0f', padding: 12, minHeight: 80, fontSize: 10, color: '#6b7280' }}>
              {monitor.outputPreview.map((line, idx) => (
                <React.Fragment key={`${line}-${idx}`}>
                  {line}
                  <br />
                </React.Fragment>
              ))}
              {currentRunId ? <span style={{ color: '#10b981' }}>{`// run: ${currentRunId}`}</span> : null}
            </div>
          </div>
        </aside>
      </section>
    </div>
  )
}

function App() {
  React.useEffect(() => {
    document.body.style.margin = '0'
    document.body.style.fontFamily = FONT_UI
    document.body.style.color = '#0f172a'

    const id = 'a2ui-global-fonts'
    if (!document.getElementById(id)) {
      const style = document.createElement('style')
      style.id = id
      style.textContent = `
        * { box-sizing: border-box; }
        h1,h2,h3,h4,h5,h6 { font-family: ${FONT_DISPLAY}; letter-spacing: 0.01em; }
        p,span,div,label,input,button,table,th,td,a { font-family: ${FONT_UI}; }
        a { color: #334155; text-decoration: none; }
        a:hover { color: #0f172a; }
      `
      document.head.appendChild(style)
    }
  }, [])

  return (
    <BrowserRouter>
      <div style={{ display: 'flex', minHeight: '100vh', background: '#f8fafc' }}>
        <IconSidebar />
        <main style={{ flex: 1, minWidth: 0 }}>
          <Routes>
            <Route path="/workspaces" element={<Page title="Workspaces" />} />
            <Route path="/workspaces/:workspaceId/sessions" element={<Page title="Sessions" />} />
            <Route path="/sessions/:sessionId/editor" element={<Page title="Editor + Preview" />} />
            <Route path="/playground" element={<Page title="Playground" />} />
            <Route path="/agent" element={<AgentPage />} />
            <Route path="/provider-pool" element={<ProviderPoolPage />} />
            <Route path="/debug" element={<DebugPage />} />
            <Route path="/settings" element={<Page title="Settings" />} />
            <Route path="*" element={<Page title="A2UI Web" />} />
          </Routes>
        </main>
      </div>
    </BrowserRouter>
  )
}

createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
