import React from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom'

function Page({ title }: { title: string }) {
  return (
    <div style={{ padding: 24 }}>
      <h2>{title}</h2>
      <p>Coming soon</p>
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
            <div><strong>{item.runId}</strong></div>
            <div>session: {item.sessionId}</div>
            <div>status: {item.status}</div>
          </button>
        ))}
      </div>

      <div style={{ border: '1px solid #ddd', borderRadius: 8, padding: 12 }}>
        <h3 style={{ marginTop: 0 }}>Run Detail</h3>
        {loading ? <p>Loading...</p> : null}
        {!loading && !detail ? <p>Select a run.</p> : null}
        {detail ? (
          <>
            <div style={{ marginBottom: 12 }}>
              <div><strong>Run:</strong> {detail.run.runId}</div>
              <div><strong>Status:</strong> {detail.run.status}</div>
              <div><strong>Started:</strong> {detail.run.startedAt}</div>
            </div>
            <h4>Steps</h4>
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
            <h4>Cost</h4>
            <div>tokens: {detail.cost.totalTokens}</div>
            <div>total: ${detail.cost.totalCost.toFixed(6)}</div>
          </>
        ) : null}
      </div>
    </div>
  )
}

function App() {
  return (
    <BrowserRouter>
      <div style={{ display: 'flex', gap: 12, padding: 12, borderBottom: '1px solid #ddd' }}>
        <Link to="/workspaces">Workspaces</Link>
        <Link to="/playground">Playground</Link>
        <Link to="/debug">Debug</Link>
        <Link to="/settings">Settings</Link>
      </div>
      <Routes>
        <Route path="/workspaces" element={<Page title="Workspaces" />} />
        <Route path="/workspaces/:workspaceId/sessions" element={<Page title="Sessions" />} />
        <Route path="/sessions/:sessionId/editor" element={<Page title="Editor + Preview" />} />
        <Route path="/playground" element={<Page title="Playground" />} />
        <Route path="/debug" element={<DebugPage />} />
        <Route path="/settings" element={<Page title="Settings" />} />
        <Route path="*" element={<Page title="A2UI Web" />} />
      </Routes>
    </BrowserRouter>
  )
}

createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
