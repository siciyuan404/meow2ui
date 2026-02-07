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

function App() {
  return (
    <BrowserRouter>
      <div style={{ display: 'flex', gap: 12, padding: 12, borderBottom: '1px solid #ddd' }}>
        <Link to="/workspaces">Workspaces</Link>
        <Link to="/playground">Playground</Link>
        <Link to="/settings">Settings</Link>
      </div>
      <Routes>
        <Route path="/workspaces" element={<Page title="Workspaces" />} />
        <Route path="/workspaces/:workspaceId/sessions" element={<Page title="Sessions" />} />
        <Route path="/sessions/:sessionId/editor" element={<Page title="Editor + Preview" />} />
        <Route path="/playground" element={<Page title="Playground" />} />
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
