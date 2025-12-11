import { Flex, Layout } from 'antd';
import { BrowserRouter, Route, Routes } from 'react-router';
import PostNote from './pages/PostNote';
import PostQuestions from './pages/PostQuestions';
import AppFooter from './layouts/Footer';
import Diaries from './pages/Diaries';
import Signup from './pages/auth/Signup';
import Signin from './pages/auth/Signin';
import Signout from './pages/auth/Signout';
import AuthedRoutes from './layouts/AuthedRoutes';
import AppHeader from './layouts/Header';

const layoutStyle: React.CSSProperties = {
  width: '100%',
  height: '100vh',
  textAlign: 'center',
};

const innerBodyStyle: React.CSSProperties = {
  width: '100%',
  maxWidth: 680,
  height: 'max-content',
  margin: 'auto',
};

function App() {
  return (
    <Layout hasSider={false} style={layoutStyle}>
      <AppHeader />
      <Flex vertical style={innerBodyStyle}>
        <BrowserRouter>
          <Routes>
            <Route path="/auth/signup" element={<Signup />} />
            <Route path="/auth/signin" element={<Signin />} />
            <Route element={<AuthedRoutes />}>
              <Route path="/" element={<PostNote />} />
              <Route path="/post" element={<PostNote />} />
              <Route path="/questions" element={<PostQuestions />} />
              <Route path="/diaries" element={<Diaries />} />
              <Route path="/auth/signout" element={<Signout />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </Flex>
      <AppFooter />
    </Layout>
  );
}

export default App;
