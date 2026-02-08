import { describe, expect, it } from 'vitest'
import type { AgentMessage, AgentMonitor, AgentRunOutput } from './agent-types'

describe('AgentRunOutput type contract', () => {
  it('has required fields', () => {
    const output: AgentRunOutput = {
      runId: 'run-1',
      versionId: 'ver-1',
      schemaJSON: '{"version":"1.0.0"}',
      repaired: false,
    }
    expect(output.runId).toBe('run-1')
    expect(output.versionId).toBe('ver-1')
    expect(output.schemaJSON).toContain('1.0.0')
    expect(output.repaired).toBe(false)
  })

  it('repaired can be true', () => {
    const output: AgentRunOutput = {
      runId: 'run-2',
      versionId: 'ver-2',
      schemaJSON: '{}',
      repaired: true,
    }
    expect(output.repaired).toBe(true)
  })
})

describe('AgentMessage type contract', () => {
  it('supports user role', () => {
    const msg: AgentMessage = {
      id: 'msg-1',
      role: 'user',
      text: 'create a dashboard',
    }
    expect(msg.role).toBe('user')
    expect(msg.text).toBe('create a dashboard')
  })

  it('supports assistant role with status', () => {
    const msg: AgentMessage = {
      id: 'msg-2',
      role: 'assistant',
      text: 'generating...',
      status: 'running',
    }
    expect(msg.role).toBe('assistant')
    expect(msg.status).toBe('running')
  })

  it('supports system role', () => {
    const msg: AgentMessage = {
      id: 'msg-3',
      role: 'system',
      text: 'session initialized',
    }
    expect(msg.role).toBe('system')
  })

  it('status is optional', () => {
    const msg: AgentMessage = {
      id: 'msg-4',
      role: 'user',
      text: 'hello',
    }
    expect(msg.status).toBeUndefined()
  })
})

describe('AgentMonitor type contract', () => {
  it('tracks token usage', () => {
    const monitor: AgentMonitor = {
      inputTokens: 100,
      outputTokens: 200,
      totalCost: 0.05,
      thinking: ['analyzing prompt', 'generating schema'],
      files: ['schema.json'],
      outputPreview: ['{', '  "type": "frame"', '}'],
    }
    expect(monitor.inputTokens + monitor.outputTokens).toBe(300)
    expect(monitor.totalCost).toBeGreaterThan(0)
    expect(monitor.thinking).toHaveLength(2)
    expect(monitor.files).toHaveLength(1)
    expect(monitor.outputPreview).toHaveLength(3)
  })

  it('handles zero state', () => {
    const monitor: AgentMonitor = {
      inputTokens: 0,
      outputTokens: 0,
      totalCost: 0,
      thinking: [],
      files: [],
      outputPreview: [],
    }
    expect(monitor.inputTokens).toBe(0)
    expect(monitor.thinking).toHaveLength(0)
  })
})
