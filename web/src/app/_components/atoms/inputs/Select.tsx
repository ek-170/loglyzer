/* eslint-disable tailwindcss/no-custom-classname */
import { Width } from '@/app/_types/tailwind/Sizing';
import { Margin, Padding } from '@/app/_types/tailwind/Spacing';
import { ComponentPropsWithRef } from 'react';
import { tv } from 'tailwind-variants';
// import { v4 as uuidv4 } from 'uuid';

export type OptionProps = {
  text: string;
} & ComponentPropsWithRef<'option'>;

export type SelectProps = {
  color: 'light' | 'dark' | 'primary';
  label?: string;
  padding?: Padding[];
  margin?: Margin[];
  width?: Width;
  optionItems: OptionProps[];
} & ComponentPropsWithRef<'select'>;

export const Select = (props: SelectProps) => {
  const { color, label, padding, margin, width, optionItems } = props;

  const paddings = padding && padding.length > 0 ? padding : ['p-0.5'];
  const margins = margin && margin.length > 0 ? margin : ['my-1.5'];
  const paddingClass = paddings.join(' ');
  const marginClass = margins.join(' ');
  const widthClass = width || 'w-fit';
  const wrapperTv = tv({
    base: `flex flex-col ${paddingClass} ${marginClass} ${widthClass}`,
  });
  return (
    <div className={wrapperTv()}>
      {label && (
        <label className="mb-2.5 block text-sm font-medium text-black-500">
          {label}
        </label>
      )}
      <select name='' className="mb-4 flex flex-col border border-gray-300 rounded focus:border-sky-500 focus:outline-none focus:ring-1 focus:ring-sky-500" data-te-select-init>
        {optionItems.length > 0 &&
          optionItems.map((optionItem) => {
            const { text, ...option } = optionItem;
            return (
              <option {...option} key={optionItem.text} className='text-md text-black-500'>
                {optionItem.text}
              </option>
            );
          })}
      </select>
    </div>
  );
};
