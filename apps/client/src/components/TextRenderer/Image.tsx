import NextImage from 'next/image';

export default function Image({
  data,
  tunes,
}: {
  data: { file: { url: string }; caption: string };
  tunes: { caption: boolean };
}) {
  return (
    <figure>
      <NextImage
        src={data.file.url}
        alt={data.caption}
        loading="lazy"
        width={1200}
        height={800}
      />
      <figcaption>{data.caption}</figcaption>
    </figure>
  );
}
