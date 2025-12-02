const DEFAULT_EXCERPT_LENGTH = 150;

type IEditorParagraphData = {
  text: string;
};
type IEditorHeadingData = {
  text: string;
  level: number;
};
type IEditorImageData = {
  file: { url: string };
  caption: string;
};
type IEditorListData = {
  items: { content: string }[];
  meta: { counterType: '' | 'numeric' };
  style: 'ordered' | 'unordered';
};
type IEditorLinkToolData = {
  link: string;
  meta: { title: string; description: string; image: { url: string } };
};

interface IEditorBlock {
  type: string;
  data:
    | IEditorHeadingData
    | IEditorParagraphData
    | IEditorImageData
    | IEditorListData
    | IEditorLinkToolData;
}

const extractTextFromBlock = (block: IEditorBlock): string => {
  switch (block.type) {
    case 'paragraph':
      return (block.data as IEditorParagraphData).text || '';
    case 'header':
      return (block.data as IEditorHeadingData).text || '';
    case 'list':
      return (
        (block.data as IEditorListData).items
          .map((item) => item.content)
          .join(' ') || ''
      );
    case 'linkTool':
      return (block.data as IEditorLinkToolData).meta.title || '';
    default:
      return '';
  }
};

const getExcerpt = (content: string, excerptLenght: number): string => {
  if (!content) {
    return '';
  }

  try {
    const { blocks, ...attrs } = JSON.parse(content || '{}');
    const excerptBlocks: IEditorBlock[] = [];
    let text = '';
    for (let i = 0; i < blocks.length; i++) {
      const content = extractTextFromBlock(blocks[i]);
      if (!content) continue;

      if (
        text.length + content.length <=
        (excerptLenght || DEFAULT_EXCERPT_LENGTH)
      ) {
        excerptBlocks.push(blocks[i]);
        text += content + ' ';
      } else {
        if (blocks[i].type === 'paragraph') {
          const remainingLength =
            (excerptLenght || DEFAULT_EXCERPT_LENGTH) - text.length;
          const paragraphData = blocks[i].data as IEditorParagraphData;
          const truncatedText = paragraphData.text.slice(0, remainingLength);
          excerptBlocks.push({
            type: 'paragraph',
            data: {
              text:
                truncatedText +
                (truncatedText.length < paragraphData.text.length ? '...' : ''),
            },
          });
        }
        // If the block is not a paragraph, we just stop adding more blocks

        break;
      }
    }
    console.log(text);

    return JSON.stringify({
      blocks: excerptBlocks,
      ...attrs,
    });
  } catch (error) {
    return '';
  }
};

export { getExcerpt };
