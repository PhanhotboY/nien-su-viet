'use client';

import { useEffect, useRef, useState } from 'react';
import type EditorJS from '@editorjs/editorjs';

import './index.css';

export default function TextEditor({
  value,
  onChange,
}: {
  value: any;
  onChange: (...args: any[]) => any;
}) {
  const isReady = useRef(false);
  const [editor, setEditor] = useState<EditorJS>();

  useEffect(() => {
    (async () => {
      if (!isReady.current) {
        // Dynamically import EditorJS and plugins only on client side
        await Promise.all([
          import('@editorjs/editorjs'),
          import('@editorjs/underline'),
          import('@editorjs/paragraph'),
          import('@editorjs/delimiter'),
          import('@editorjs/warning'),
          import('@editorjs/header'),
          import('@editorjs/table'),
          import('@editorjs/image'),
          import('@editorjs/quote'),
          import('@editorjs/list'),
          // @ts-ignore - no types available
          import('@editorjs/raw'),
          // @ts-ignore - no types available
          import('@editorjs/link'),
          // @ts-ignore - no types available
          import('@editorjs/embed'),
          // @ts-ignore - no types available
          import('@editorjs/marker'),
          // @ts-ignore - no types available
          import('editorjs-text-alignment-blocktune'),
        ]).then(
          ([
            { default: EditorJS },
            { default: Underline },
            { default: Paragraph },
            { default: Delimiter },
            { default: Warning },
            { default: Header },
            { default: Table },
            { default: Image },
            { default: Quote },
            { default: List },
            { default: Raw },
            { default: Link },
            { default: Embed },
            { default: Marker },
            { default: AlignmentTuneTool },
          ]) => {
            const tools = {
              list: {
                class: List,
                inlineToolbar: true,
              },
              header: {
                class: Header,
                tunes: ['textAlign'],
              },
              paragraph: {
                class: Paragraph,
                tunes: ['textAlign'],
              },
              image: {
                class: Image,
                config: {
                  endpoints: {
                    byFile: '/api/images/upload',
                    byUrl: '/api/images/fetchUrl',
                  },
                  field: 'img',
                },
              },
              linkTool: {
                class: Link,
                config: {
                  endpoint: '/api/fetchUrl',
                },
              },
              html: Raw,
              embed: Embed,
              table: {
                class: Table,
                inlineToolbar: true,
                config: {
                  maxRows: 5,
                  maxCols: 5,
                },
              },
              quote: Quote,
              marker: Marker,
              warning: Warning,
              underline: Underline,
              delimiter: Delimiter,
              textAlign: {
                class: AlignmentTuneTool,
                config: {
                  default: 'left',
                },
              },
            };

            // prevent initializing editor more than once
            if (!isReady.current) {
              const editor = new EditorJS({
                holder: 'editorjs',
                // @ts-ignore
                tools,
                data: value && JSON.parse(value),
                onChange: (api, e) => {
                  api.saver.save().then((outputData: any) => {
                    onChange(JSON.stringify(outputData));
                  });
                },
              });

              setEditor(editor);
            }
            isReady.current = true;
          },
        );
      }
    })();

    return () => {
      if (editor && editor.destroy) {
        editor.destroy();
      }
    };
  }, []);

  return <div id="editorjs" className="border rounded mt-2"></div>;
}
