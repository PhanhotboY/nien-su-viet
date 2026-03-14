import NextImage from 'next/image';

export default function Image({
  data,
  tunes,
}: {
  data: { file: { url: string }; caption: string };
  tunes?: { caption: boolean };
}) {
  return (
    <figure className="max-w-[1200px] max-h-[800px] overflow-hidden">
      <NextImage
        src={data.file.url}
        alt={data.caption}
        loading="lazy"
        width={1200}
        height={800}
        // className="w-full h-full object-cover object-center"
      />
      {tunes?.caption && <figcaption>{data.caption}</figcaption>}
    </figure>
  );
}
