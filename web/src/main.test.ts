import { describe, expect, it } from 'vitest'
import type { AgentMonitor } from './agent-types'

describe('web smoke', () => {
  it('works', () => {
    expect(true).toBe(true)
  })

  it('agent monitor type contract is stable', () => {
    const monitor: AgentMonitor = {
      inputTokens: 1,
      outputTokens: 2,
      totalCost: 0.1,
      thinking: ['x'],
      files: ['a.json'],
      outputPreview: ['{', '}'],
    }
    expect(monitor.inputTokens + monitor.outputTokens).toBe(3)
    expect(monitor.files.length).toBeGreaterThan(0)
  })
})
