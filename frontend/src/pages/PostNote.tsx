import { LoadingOutlined } from '@ant-design/icons';
import { Button, Card, Flex, Form, Input, Segmented, Typography } from 'antd';
import { useForm } from 'antd/es/form/Form';
import { useCallback, useState } from 'react';
import useSWR from 'swr';
import { authorizeFetcherJson } from '../lib/fetcher';
import type { Question } from '../models/question';
import useAuthStorage from '../lib/useAuthStorage';

const { Item: FormItem } = Form;
const { TextArea } = Input;
const { Title, Link } = Typography;

type PostForm = {
  theme: number;
  note: string;
};

export default function PostNote() {
  const { token } = useAuthStorage();
  const {
    data: questions,
    isLoading,
    mutate,
  } = useSWR<Question[]>('/v1/api/questions', (url: string) =>
    authorizeFetcherJson(url, token || ''),
  );
  const [isSending, setIsSending] = useState(false);
  const [form] = useForm<PostForm>();

  const handleSubmit = useCallback(
    (values: PostForm) => {
      setIsSending(true);
      fetch('/v1/api/diaries', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          questionId: values.theme,
          note: values.note,
          userId: 1,
        }),
      })
        .then(() => setIsSending(false))
        .then(() => form.resetFields());
    },
    [form, token],
  );

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
          <Form
            form={form}
            layout="vertical"
            onFinish={handleSubmit}
            initialValues={{ theme: 1 }}
          >
            {isLoading && <LoadingOutlined />}
            {questions && (
              <FormItem
                name={'theme'}
                label="今日のテーマ"
                rules={[{ required: true }]}
              >
                <Segmented
                  options={questions.map((qItem) => {
                    return { label: qItem.qtext, value: qItem.id };
                  })}
                  size="large"
                />
              </FormItem>
            )}
            <Link style={{ textAlign: 'right' }} onClick={() => mutate()}>
              選び直す？
            </Link>
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
            <FormItem>
              <Button type="primary" htmlType="submit" loading={isSending}>
                記録！
              </Button>
            </FormItem>
          </Form>
        </Flex>
      </Card>
    </Flex>
  );
}
