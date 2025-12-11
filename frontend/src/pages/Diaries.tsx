import useSWR from 'swr';
import { authorizeFetcherJson } from '../lib/fetcher';
import type { Diary } from '../models/diary';
import { Card, Flex, Typography } from 'antd';
import useAuthStorage from '../lib/useAuthStorage';
import { useCallback, useState } from 'react';
import { DeleteOutlined } from '@ant-design/icons';
const { Title, Text } = Typography;

export default function Diaries() {
  const { token } = useAuthStorage();
  const [selectedId, setSelectedId] = useState<number | undefined>();
  const {
    data: daiaries,
    // isLoading,
    mutate,
  } = useSWR<Diary[]>('/v1/api/diaries', (url: string) =>
    authorizeFetcherJson(url, token || ''),
  );

  const deleteOnClick = useCallback(
    (id: number) => {
      fetch(`/v1/api/diaries/${id}`, {
        method: 'DELETE',
        headers: { Authorization: `Bearer ${token}` },
      })
        .then(() => mutate())
        .catch((err) => console.error(err));
    },
    [token, mutate],
  );
  return (
    <Flex vertical gap={'middle'}>
      <Flex vertical>
        <Title level={2}>あなたの日記</Title>
        <Typography>あなたが過去に投稿した一行日記を振り返る。</Typography>
      </Flex>
      <Flex vertical gap={'small'}>
        {daiaries?.map((it, idx) => (
          <DiaryCard
            diary={it}
            key={idx}
            selected={selectedId === it.id}
            onClick={setSelectedId}
            deleteOnClick={deleteOnClick}
          />
        ))}
      </Flex>
    </Flex>
  );
}

type DiaryCardProps = {
  diary: Diary;
  selected: boolean;
  onClick: (id: number | undefined) => void;
  deleteOnClick: (id: number) => void;
};

function DiaryCard({
  diary,
  selected,
  onClick,
  deleteOnClick,
}: DiaryCardProps) {
  const actions = selected
    ? [
        <DeleteOutlined
          key={'delete'}
          onClick={() => deleteOnClick(diary.id)}
        />,
      ]
    : [];
  const selectedStyle = selected
    ? {
        borderColor: '#696FC7',
        boxShadow: '0 2px 8px #A7AAE1',
      }
    : {};
  return (
    <Card
      style={{
        width: '100%',
        textAlign: 'left',
        ...selectedStyle,
      }}
      actions={actions}
      onClick={() => onClick(selected ? undefined : diary.id)}
    >
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
