import { Flex } from 'antd';
import { BrowserRouter, Route, Routes } from 'react-router';
import PostNote from './pages/PostNote';

const layoutStyle: React.CSSProperties = {
  width: '100%',
  maxWidth: 980,
  height: 'max-content',
  minHeight: '100%',
  margin: '0 auto',
};

function App() {
  return (
    <Flex vertical style={layoutStyle}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<PostNote />} />
        </Routes>
      </BrowserRouter>
    </Flex>
  );
}

export default App;
