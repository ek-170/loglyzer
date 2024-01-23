/* eslint-disable tailwindcss/no-custom-classname */
import { Width } from '@/app/_types/tailwind/Sizing';
import { Margin, Padding } from '@/app/_types/tailwind/Spacing';
import { ComponentPropsWithRef, useState } from 'react';
import { tv } from 'tailwind-variants';

export type FileProps = {
  color: 'light' | 'dark' | 'primary';
  label?: string;
  padding?: Padding[];
  margin?: Margin[];
  width?: Width;
  inputElementProps: Omit<ComponentPropsWithRef<'input'>, 'type' | 'multiple'>;
};

export const File = (props: FileProps) => {
  const [fileName, setFileName] = useState<string>('選択されていません');

  const { color, padding, margin, width, label, ...fileElementProps } = props;

  const paddings = padding && padding.length > 0 ? padding : ['p-0.5'];
  const margins = margin && margin.length > 0 ? margin : ['my-1.5'];
  const paddingClass = paddings.join(' ');
  const marginClass = margins.join(' ');
  const widthClass = width || 'w-fit';
  const wrapperTv = tv({
    base: `${paddingClass} ${marginClass} ${widthClass}`,
  });
  const handleOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.currentTarget.files;
    if (!files || files?.length === 0) {
      return;
    }
    setFileName(files?.item(0)?.name ?? "選択されていません");
  };

  return (
    <div className={wrapperTv()}>
      {label && (
        <label className="mb-2.5 block text-sm font-medium text-black-500">
          {label}
        </label>
      )}
      <label className="mr-3 mb-1.5 inlineblock rounded bg-blue-500 p-2 hover:bg-blue-300">
        <input
          {...fileElementProps}
          type="file"
          className={'hidden'}
          onChange={handleOnChange}
        />
        <span className={'text-sm text-black-50'}>
          ファイルを選択してください
        </span>
      </label>
      <span className={'inline-block w-48 text-sm text-black-500'}>
        {fileName}
      </span>
    </div>
  );
};
