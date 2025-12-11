import type { Question } from '../models/question';
import { Card, Masonry, Typography } from 'antd';
const { Text } = Typography;

export default function Questions({ questions }: { questions: Question[] }) {
  return (
    <Masonry
      gutter={16}
      items={questions?.map((it, key) => {
        return { key: key, data: it };
      })}
      itemRender={({ data, index }) => (
        <QuestionCard question={data} key={index} />
      )}
    />
  );
}

function QuestionCard({ question }: { question: Question }) {
  return (
    <Card>
      <Text>{question.qtext}</Text>
    </Card>
  );
}
