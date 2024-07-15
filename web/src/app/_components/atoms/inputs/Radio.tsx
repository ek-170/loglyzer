/* eslint-disable tailwindcss/no-custom-classname */
import { Width } from '@/app/_types/tailwind/Sizing';
import { Margin, Padding } from '@/app/_types/tailwind/Spacing';
import { ComponentPropsWithRef } from 'react';
import { tv } from 'tailwind-variants';
import { v4 as uuidv4 } from 'uuid';

type RadioItemProps = {
  text: string;
} & Omit<ComponentPropsWithRef<'input'>, 'type' | 'name'>;

export type RadioProps = {
  color: 'light' | 'dark' | 'primary';
  label?: string;
  name: string;
  padding?: Padding[];
  margin?: Margin[];
  width?: Width;
  radioItems: RadioItemProps[];
};

export const Radio = (props: RadioProps) => {
  const { color, label, name, padding, margin, width, radioItems } = props;

  const paddings = padding && padding.length > 0 ? padding : ['p-0.5'];
  const margins = margin && margin.length > 0 ? margin : ['my-1.5'];
  const paddingClass = paddings.join(' ');
  const marginClass = margins.join(' ');
  const widthClass = width || 'w-fit';
  const wrapperTv = tv({
    base: `flex flex-col ${paddingClass} ${marginClass} ${widthClass}`,
  });
  const textTv = tv({
    base: `s-2 pl-2 text-sm font-medium`,
    variants: {
      color: {
        light: 'text-black-600',
        dark: 'text-black-50',
        primary: 'text-primary-500',
      },
    },
  });
  return (
    <div className={wrapperTv()}>
      {label && (
        <label className="mb-2.5 block text-sm font-medium text-black-500">
          {label}
        </label>
      )}
      <div className="mb-4 flex flex-col">
        {radioItems.length > 0 &&
          radioItems.map((radioItem) => {
            const inputId = uuidv4();
            const { text, ...props } = radioItem;
            return (
              <div key={text}>
                <input {...props} id={inputId} type="radio" name={name} />
                <label htmlFor={inputId} className={textTv({ color: color })}>
                  {text}
                </label>
              </div>
            );
          })}
      </div>
    </div>
  );
};
