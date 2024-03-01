import { createBrowserRouter, RouteObject } from 'react-router-dom';

import NotFound from '@/404';
import InternalServerError from '@/500';
import Home from '@/page/Home';
import Layout from '@/page/Layout';

const routes: RouteObject[] = [
  {
    path: '/',
    element: <Layout />,
    errorElement: <InternalServerError />,
    children: [
      {
        index: true,
        element: <Home />,
      },
    ],
  },
  {
    path: '*',
    element: <NotFound />,
  },
];

const router = createBrowserRouter(routes);

export default router;
