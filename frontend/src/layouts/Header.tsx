import { Flex, Layout, Typography } from 'antd';
import useAuthStorage from '../lib/useAuthStorage';
const { Header } = Layout;
const { Title, Link } = Typography;
export default function AppHeader() {
  const { token } = useAuthStorage();

  return (
    <Header style={{ backgroundColor: '#9adc79ff' }}>
      <Flex align="center">
        <Flex flex={1}>
          <Title
            level={4}
            style={{ margin: '1em', fontFamily: 'Delius, cursive' }}
          >
            One-Line Diary
          </Title>
        </Flex>
        {token ? (
          <>
            <Link href="/auth/signout">サインアウト</Link>
          </>
        ) : (
          <>
            <Link href="/auth/signin">ログイン</Link>
            <Link href="/auth/signup">登録</Link>
          </>
        )}
      </Flex>
    </Header>
  );
}
