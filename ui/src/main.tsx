import './normalize.css';

import { App } from 'antd';
import React from 'react';
import ReactDOM from 'react-dom/client';
import { RouterProvider } from 'react-router-dom';

import { ThemeProvider } from './provider/theme';
import router from './router';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ThemeProvider>
      <App>
        <RouterProvider router={router} />
      </App>
    </ThemeProvider>
  </React.StrictMode>,
);
