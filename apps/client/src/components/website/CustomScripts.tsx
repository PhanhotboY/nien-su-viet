'use client';

import { useEffect } from 'react';

interface CustomScriptsProps {
  headScripts?: string;
  bodyScripts?: string;
}

export const CustomScripts = ({
  headScripts,
  bodyScripts,
}: CustomScriptsProps) => {
  useEffect(() => {
    // Inject head scripts
    if (headScripts) {
      const headContainer = document.createElement('div');
      headContainer.innerHTML = headScripts;
      const scripts = headContainer.querySelectorAll('script');
      const styles = headContainer.querySelectorAll('style, link');
      const metas = headContainer.querySelectorAll('meta');

      // Append meta tags
      metas.forEach((meta) => {
        const clonedMeta = meta.cloneNode(true) as HTMLMetaElement;
        document.head.appendChild(clonedMeta);
      });

      // Append styles and link tags
      styles.forEach((style) => {
        const clonedStyle = style.cloneNode(true) as HTMLElement;
        document.head.appendChild(clonedStyle);
      });

      // Append and execute scripts
      scripts.forEach((script) => {
        const newScript = document.createElement('script');
        if (script.src) {
          newScript.src = script.src;
        } else {
          newScript.textContent = script.textContent;
        }
        if (script.async) newScript.async = true;
        if (script.defer) newScript.defer = true;
        Array.from(script.attributes).forEach((attr) => {
          if (attr.name !== 'src') {
            newScript.setAttribute(attr.name, attr.value);
          }
        });
        document.head.appendChild(newScript);
      });
    }

    // Inject body scripts
    if (bodyScripts) {
      const bodyContainer = document.createElement('div');
      bodyContainer.innerHTML = bodyScripts;
      const scripts = bodyContainer.querySelectorAll('script');
      const otherElements = bodyContainer.querySelectorAll('*:not(script)');

      // Append non-script elements
      otherElements.forEach((element) => {
        const clonedElement = element.cloneNode(true) as HTMLElement;
        document.body.appendChild(clonedElement);
      });

      // Append and execute scripts
      scripts.forEach((script) => {
        const newScript = document.createElement('script');
        if (script.src) {
          newScript.src = script.src;
        } else {
          newScript.textContent = script.textContent;
        }
        if (script.async) newScript.async = true;
        if (script.defer) newScript.defer = true;
        Array.from(script.attributes).forEach((attr) => {
          if (attr.name !== 'src') {
            newScript.setAttribute(attr.name, attr.value);
          }
        });
        document.body.appendChild(newScript);
      });
    }

    // Cleanup function to remove injected scripts on unmount
    return () => {
      // Note: We don't remove scripts on unmount as they may be needed for analytics, etc.
      // If you need to remove them, you'd need to track them in a ref
    };
  }, [headScripts, bodyScripts]);

  return null;
};
