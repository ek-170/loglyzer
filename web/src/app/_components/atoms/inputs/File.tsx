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

  const [fileNames, setFileNames] = useState<string[]>(["選択されていません"]);

  const {
    color,
    padding,
    margin,
    width,
    label,
    ...fileElementProps
  } = props;

  const paddings = padding && padding.length > 0 ? padding : ['p-0.5'];
  const margins = margin && margin.length > 0 ? margin : ['my-1.5'];
  const paddingClass = paddings.join(' ');
  const marginClass = margins.join(' ');
  const widthClass = width || 'w-fit';
  const wrapperTv = tv({
    base: `flex-col ${paddingClass} ${marginClass} ${widthClass}`,
  });
  const handleOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.currentTarget.files;
    if (!files || files?.length === 0) return;
    console.log(files)
    const names: string[] = [];
    for (let i = 0; i <= files.length; i++){
      const file = files.item(i);
      if (file?.name) names.push(file?.name);
    }
    setFileNames(names);
  }

  return (
    <div className={wrapperTv()}>
      {label && (
        <label className="mb-2.5 block text-sm font-medium text-black-500">
          {label}
        </label>
      )}
      <label className='bg-blue-500 inline-block p-2 mr-3 rounded hover:bg-blue-300'>
        <input {...fileElementProps} type="file" className={"hidden"} onChange={handleOnChange} multiple/>
        <span className={"text-black-50 text-sm"}>
          ファイルを選択してください
        </span>
      </label>
      {
        fileNames.map((f) => {
          return (
            <>
              <span key={f} className={"text-black-500 text-sm"}>
                {f}
              </span>
              <br/>
            </>
          );
        })
      }
    </div>
  );
};
