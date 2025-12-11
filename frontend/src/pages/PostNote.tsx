import { LoadingOutlined } from '@ant-design/icons';
import { Button, Card, Flex, Form, Input, Radio, Typography } from 'antd';
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
      console.log(values.theme);
      if (!values.theme) return;
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
          <Form form={form} layout="vertical" onFinish={handleSubmit}>
            {isLoading && <LoadingOutlined />}

            <FormItem
              name={'theme'}
              label="今日のテーマ"
              rules={[
                { required: true, message: '回答する質問を選びましょう。' },
              ]}
            >
              <Radio.Group>
                {questions?.map((it, key) => (
                  <Radio.Button value={it.id} key={key}>
                    {it.qtext}
                  </Radio.Button>
                ))}
              </Radio.Group>
            </FormItem>

            <Link style={{ textAlign: 'right' }} onClick={() => mutate()}>
              選び直す？
            </Link>
            <FormItem
              name={'note'}
              label={'今日の記録'}
              rules={[
                { required: true, message: '「ひとこと」を入力しましょう' },
              ]}
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
