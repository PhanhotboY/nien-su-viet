import NextImage from 'next/image';

export default function Image({
  data,
  tunes,
}: {
  data: { file: { url: string }; caption: string };
  tunes: { caption: boolean };
}) {
  return (
    <div>
      <NextImage src={data.file.url} alt={data.caption} loading="lazy" />
    </div>
  );
}
