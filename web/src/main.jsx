import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import { HashRouter, Routes, Route } from "react-router-dom";
import HomePage from './pages/HomePage';
import TagPage from './pages/TagPage';
import TokenPage from './pages/TokenPage';
import LoginPage from './pages/LoginPage';

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <HashRouter>
      <UserProvider>
        <Routes>
          <Route path="/" element={<App />}>
            <Route path="home" element={<RequireAuth><HomePage/></RequireAuth>} />
            <Route path="tags" element={<RequireAuth><TagPage/></RequireAuth>} />
            <Route path="tokens" element={<RequireAuth><TokenPage/></RequireAuth>} />
            <Route path="/" element={<LoginPage />} />
          </Route>
        </Routes>
      </UserProvider>
    </HashRouter>
  </React.StrictMode>,
)
