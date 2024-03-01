import { Layout, theme } from 'antd';
import { Outlet } from 'react-router-dom';

import Header from './Header';

const { Header: LayoutHeader, Content } = Layout;

export default function Basic() {
  const { token } = theme.useToken();

  return (
    <Layout>
      <LayoutHeader
        style={{
          position: 'sticky',
          top: 0,
          zIndex: 1,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          backgroundColor: token.colorBgContainer,
        }}
      >
        <Header />
      </LayoutHeader>
      <Content
        style={{
          display: 'flex',
          justifyContent: 'center',
          height: 'calc(100vh - 64px)',
          overflow: 'auto',
        }}
      >
        <div
          style={{ width: '100%', maxWidth: 1200, paddingTop: token.padding }}
        >
          <Outlet />
        </div>
      </Content>
    </Layout>
  );
}
