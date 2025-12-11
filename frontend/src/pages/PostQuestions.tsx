import { Button, Card, Flex, Form, Input, Typography } from 'antd';
import { useForm } from 'antd/es/form/Form';
import { useCallback, useState } from 'react';
import useAuthStorage from '../lib/useAuthStorage';
import Questions from '../components/Questions';
import type { Question } from '../models/question';
import useSWR from 'swr';
import { authorizeFetcherJson } from '../lib/fetcher';
const { Item: FormItem } = Form;
const { Title } = Typography;

type QuestionForm = {
  qtext: string;
};

export default function PostQuestions() {
  const [form] = useForm<QuestionForm>();
  const [isSending, setIsSending] = useState(false);
  const { token } = useAuthStorage();
  const { data: questions, mutate } = useSWR<Question[]>(
    '/v1/api/questions?mine=true',
    (url: string) => authorizeFetcherJson(url, token || ''),
  );

  const handleSubmit = useCallback(
    (values: QuestionForm) => {
      fetch('/v1/api/questions', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(values),
      })
        .then(() => setIsSending(false))
        .then(() => form.resetFields())
        .then(() => mutate());
    },
    [form, token, mutate],
  );

  return (
    <Flex vertical gap={'large'}>
      <Card style={{ width: '100%', maxWidth: 680, margin: 'auto' }}>
        <Flex vertical>
          <Title level={3}>出題する質問を追加しよう。</Title>
          <Form form={form} onFinish={handleSubmit} layout="vertical">
            <FormItem
              name={'qtext'}
              label={'質問テーマ'}
              rules={[{ required: true, message: '質問を入力してください！' }]}
            >
              <Input />
            </FormItem>
            <FormItem>
              <Button type="primary" htmlType="submit" loading={isSending}>
                登録
              </Button>
            </FormItem>
          </Form>
        </Flex>
      </Card>
      <Title level={4}>あなたが投稿した質問</Title>
      {questions && <Questions questions={questions} />}
    </Flex>
  );
}
