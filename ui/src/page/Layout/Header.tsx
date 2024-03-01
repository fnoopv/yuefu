import {
  IconBell,
  IconDeviceDesktop,
  IconLogout,
  IconMoonStars,
  IconSun,
  IconUserCircle,
} from '@tabler/icons-react';
import {
  Avatar,
  Badge,
  Button,
  Dropdown,
  Empty,
  Menu,
  Popover,
  Space,
  Tabs,
  theme as antdTheme,
} from 'antd';
import { useEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

import { useTheme } from '@/hook/theme';

export default function Header() {
  const { token } = antdTheme.useToken();
  const location = useLocation();
  const navigate = useNavigate();
  const [activeKey, setActiveKey] = useState(location.pathname);
  const { theme, setTheme } = useTheme();

  useEffect(() => {
    if (location.pathname !== activeKey) {
      setActiveKey(location.pathname);
    }
  }, [activeKey, location]);

  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        justifyContent: 'space-between',
        maxWidth: 1440,
      }}
    >
      <div style={{ flex: 1, display: 'flex' }}>
        <div
          style={{
            display: 'flex',
            alignItems: 'center',
          }}
        >
          <img
            src='/logo.svg'
            height={48}
            onClick={() => navigate('/')}
            style={{ cursor: 'pointer' }}
          />
        </div>
        <Menu
          mode='horizontal'
          items={[
            { label: '发现', key: '/', onClick: () => navigate('/') },
            { label: '我的', key: '/my', onClick: () => navigate('/my') },
            {
              label: '管理',
              key: '/manage',
              onClick: () => navigate('/manager'),
            },
            {
              label: '配置',
              key: '/config',
              onClick: () => navigate('/config'),
            },
          ]}
          selectedKeys={[activeKey]}
          style={{
            flex: 1,
            backgroundColor: token.colorBgContainer,
            border: 'none',
            marginLeft: token.margin,
          }}
        />
      </div>
      <Space direction='horizontal' style={{ color: token.colorText }}>
        <Dropdown
          menu={{
            items: [
              {
                label: (
                  <div
                    style={{
                      display: 'flex',
                      alignItems: 'center',
                    }}
                  >
                    <IconDeviceDesktop size='1rem' />
                    <span style={{ marginLeft: 4 }}>系统</span>
                  </div>
                ),
                key: 'system',
                onClick: () => setTheme('system'),
              },
              {
                label: (
                  <div
                    style={{
                      display: 'flex',
                      alignItems: 'center',
                    }}
                  >
                    <IconSun size='1rem' />
                    <span style={{ marginLeft: 4 }}>明亮</span>
                  </div>
                ),
                key: 'light',
                onClick: () => setTheme('light'),
              },
              {
                label: (
                  <div
                    style={{
                      display: 'flex',
                      alignItems: 'center',
                    }}
                  >
                    <IconMoonStars size='1rem' />
                    <span style={{ marginLeft: 8 }}>黑暗</span>
                  </div>
                ),
                key: 'dark',
                onClick: () => setTheme('dark'),
              },
            ],
          }}
          overlayStyle={{ paddingTop: 16 }}
        >
          <Button
            type='text'
            size='large'
            style={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              padding: 8,
            }}
          >
            {theme === 'dark' ? <IconSun /> : <IconMoonStars />}
          </Button>
        </Dropdown>
        <Popover
          content={
            <div style={{ width: 240, height: 320 }}>
              <Tabs
                size='small'
                centered
                items={[
                  { label: '系统通知', key: 'system', children: <Empty /> },
                  { label: '用户任务', key: 'user' },
                ]}
              />
            </div>
          }
          overlayStyle={{ paddingTop: 8 }}
        >
          {/**
           * // TODO: 有未读通知时填充
           */}
          <Button
            type='text'
            size='large'
            style={{
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              padding: 8,
            }}
          >
            <Badge size='small' count={0}>
              <IconBell />
            </Badge>
          </Button>
        </Popover>
        <Dropdown
          menu={{
            items: [
              {
                label: (
                  <div
                    style={{
                      display: 'flex',
                      alignItems: 'center',
                    }}
                  >
                    <IconUserCircle size='1rem' />
                    <span style={{ marginLeft: 4 }}>个人信息</span>
                  </div>
                ),
                key: 'profile',
              },
              {
                label: (
                  <div
                    style={{
                      display: 'flex',
                      alignItems: 'center',
                    }}
                  >
                    <IconLogout size='1rem' />
                    <span style={{ marginLeft: 4 }}>退出登录</span>
                  </div>
                ),
                key: 'logout',
              },
            ],
          }}
          overlayStyle={{ paddingTop: 16 }}
        >
          <div style={{ display: 'flex' }}>
            <Button
              type='text'
              size='large'
              style={{
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
                padding: 8,
              }}
            >
              <Avatar src='https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png' />
              <span style={{ marginLeft: 8 }}>张三</span>
            </Button>
          </div>
        </Dropdown>
      </Space>
    </div>
  );
}
