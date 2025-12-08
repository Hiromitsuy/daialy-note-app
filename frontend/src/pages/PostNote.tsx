import { LoadingOutlined } from '@ant-design/icons';
import { Button, Card, Flex, Form, Input, Segmented, Typography } from 'antd';
import { useForm } from 'antd/es/form/Form';
import useSWR from 'swr';
const { Item: FormItem } = Form;
const { TextArea } = Input;
const { Title, Link } = Typography;

const mockData = [
  {
    id: 1,
    qtext: '今の気分は？',
  },
  {
    id: 2,
    qtext: '今日の夕食は？',
  },
  {
    id: 3,
    qtext: '〇〇でいい感じ！',
  },
];

const mockFetch = () => {
  return new Promise<typeof mockData>((resolve) => {
    setTimeout(() => resolve(mockData), 1000);
  });
};

export default function PostNote() {
  const {
    data: questions,
    isLoading,
    mutate,
  } = useSWR<typeof mockData>('/api/question', mockFetch);
  const [form] = useForm();

  const getDateString = () => {
    const now = new Date();
    const dayStringArr = ['日', '月', '火', '水', '木', '金', '土'];
    return `${now.getFullYear()} 年 ${
      now.getMonth() + 1
    } 月 ${now.getDate()} 日 （${dayStringArr[now.getDay()]}）`;
  };

  return (
    <Flex vertical gap={'large'}>
      <Title level={1}>{getDateString()}</Title>
      <Card>
        <Flex gap={'middle'} vertical>
          <Title level={3}>今日の記録を書き留めよう。</Title>
          <Form form={form} layout="vertical">
            <FormItem
              name={'theme'}
              label="今日のテーマ"
              rules={[{ required: true }]}
            >
              <Flex vertical gap={'small'}>
                {isLoading && <LoadingOutlined />}
                {questions && (
                  <Segmented
                    options={questions.map((qItem) => qItem.qtext)}
                    size="large"
                  />
                )}
                <Link style={{ textAlign: 'right' }} onClick={() => mutate()}>
                  選び直す？
                </Link>
              </Flex>
            </FormItem>
            <FormItem
              name={'note'}
              label={'今日の記録'}
              rules={[{ required: true }]}
            >
              <TextArea
                rows={2}
                maxLength={255}
                showCount
                placeholder="回答を書き留める..."
              />
            </FormItem>
            <FormItem name={'submit'}>
              <Button type="primary" htmlType="submit">
                記録！
              </Button>
            </FormItem>
          </Form>
        </Flex>
      </Card>
    </Flex>
  );
}
