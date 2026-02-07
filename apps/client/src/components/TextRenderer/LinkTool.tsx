import { useEffect } from 'react';
import './LinkTool.css';

export default function LinkTool({
  data,
}: {
  data: {
    link: string;
    meta: { title: string; description: string; image: { url: string } };
  };
}) {
  useEffect(() => {
    (async () => {
      // Dynamically import EditorJS and plugins only on client side
      // @ts-ignore - no types available
      await import('@editorjs/link');
    })();
  }, []);

  return (
    <article className="cdx-block mt-4">
      <div className="link-tool">
        <a
          className="link-tool__content link-tool__content--rendered"
          target="_blank"
          rel="nofollow noindex noreferrer"
          href={data.link}
          data-empty="false"
        >
          <section>
            <div
              className="link-tool__image"
              style={{ backgroundImage: `url("${data.meta.image.url}")` }}
            ></div>
            <h3 className="link-tool__title">{data.meta.title}</h3>
          </section>
          <p className="link-tool__description">{data.meta.description}</p>
          <footer className="link-tool__anchor">
            {data.link.replace(/https?:\/\//gm, '').replace(/\/.*/gm, '')}
          </footer>
        </a>
      </div>
    </article>
  );
}
