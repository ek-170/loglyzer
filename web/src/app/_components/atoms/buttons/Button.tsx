import { Padding, Width } from '@/app/_types/tailwind/Sizing';
import { ComponentPropsWithRef, ReactNode } from 'react';
import { tv } from 'tailwind-variants';

/**
 * padding and width should be written in the way the class name of tailwind is written
 */
export type ButtonProps = {
  color: 'positive' | 'negative' | 'danger' | 'primary';
  padding?: Padding[];
  width?: Width;
  buttonElementProps: ComponentPropsWithRef<'button'>;
  children?: ReactNode;
};

export const Button = (props: ButtonProps) => {
  const paddings =
    props.padding && props.padding.length > 0
      ? props.padding
      : ['px-3', 'py-1.5'];
  const paddingClass = paddings.join(' ');
  const width = props.width || 'w-fit';
  // eslint-disable-next-line tailwindcss/no-custom-classname
  const button = tv({
    base: `flex justify-center rounded-md ${paddingClass} ${width} shadow-sm focus-visible:outline`,
    variants: {
      color: {
        positive: 'bg-sky-500 hover:bg-sky-300 disabled:bg-sky-100',
        negative: 'bg-black-400 hover:bg-black-500 disabled:bg-black-200',
        danger: 'bg-red-500 hover:bg-red-400 disabled:bg-red-200',
        primary: 'bg-primary-500 hover:bg-primary-400 disabled:bg-primary-200',
      },
    },
  });
  return (
    <>
      <button
        {...props.buttonElementProps}
        className={button({ color: props.color })}
      >
        {props.children}
      </button>
    </>
  );
};
