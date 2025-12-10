import useSWR from 'swr';
import { authorizeFetcherJson } from '../lib/fetcher';
import type { Diary } from '../models/diary';
import { Card, Flex, Typography } from 'antd';
import useAuthStorage from '../lib/useAuthStorage';
const { Title, Text } = Typography;

export default function Diaries() {
  const { token } = useAuthStorage();
  const {
    data: daiaries,
    // isLoading,
    // mutate,
  } = useSWR<Diary[]>('/v1/api/diaries', () =>
    authorizeFetcherJson('/v1/api/diaries', token || ''),
  );

  console.log(daiaries);
  return (
    <Flex vertical gap={'middle'}>
      <Flex vertical>
        <Title level={2}>あなたの日記</Title>
        <Typography>あなたが過去に投稿した一行日記を振り返る。</Typography>
      </Flex>
      <Flex vertical gap={'small'}>
        {daiaries?.map((it, idx) => (
          <DiaryCard diary={it} key={idx} />
        ))}
      </Flex>
    </Flex>
  );
}

function DiaryCard({ diary }: { diary: Diary }) {
  return (
    <Card style={{ width: '100%', textAlign: 'left' }}>
      <Flex vertical gap={'small'}>
        <Typography>{formatDatetime(diary.createdAt)}</Typography>
        <Flex gap={'middle'}>
          <Text type="secondary">{diary.question.qtext}</Text>
          <Typography>{diary.note}</Typography>
        </Flex>
      </Flex>
    </Card>
  );
}

function formatDatetime(date: Date) {
  const tmpDate = new Date(date);
  return (
    `${tmpDate.getFullYear()}年${
      tmpDate.getMonth() + 1
    }月${tmpDate.getDate()}日` +
    ` ${tmpDate.getHours()}時${tmpDate.getMinutes()}分 に投稿`
  );
}
