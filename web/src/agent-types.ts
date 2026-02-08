export type AgentRunOutput = {
  runId: string
  versionId: string
  schemaJSON: string
  repaired: boolean
}

export type AgentMessage = {
  id: string
  role: 'user' | 'assistant' | 'system'
  text: string
  status?: 'running' | 'success' | 'error'
}

export type AgentMonitor = {
  inputTokens: number
  outputTokens: number
  totalCost: number
  thinking: string[]
  files: string[]
  outputPreview: string[]
}
