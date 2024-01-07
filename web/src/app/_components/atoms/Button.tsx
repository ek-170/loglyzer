import { Padding, Width } from '@/app/_types/tailwind/Sizing';
import { ComponentPropsWithRef } from 'react';
import { ReactNode } from 'react';
import { tv } from 'tailwind-variants';

/**
 * padding and width should be written in the way the class name of tailwind is written
 */
type ButtonProps = {
  color: 'positive' | 'negative' | 'danger';
  padding?: Padding[];
  width?: Width;
  children?: ReactNode;
  buttonElementProps: ComponentPropsWithRef<'button'>;
};

export default function Button(props: ButtonProps) {
  const paddings =
    props.padding && props.padding.length > 0
      ? props.padding
      : ['px-3', 'py-1.5'];
  const paddingClass = paddings.join(' ');
  const width = props.width || 'w-20';
  const button = tv({
    base: `flex justify-center rounded-md ${paddingClass} ${width} shadow-sm focus-visible:outline`,
    variants: {
      color: {
        positive: 'bg-sky-500 hover:bg-sky-300 disabled:bg-sky-100',
        negative: 'bg-zinc-400 hover:bg-zinc-500 disabled:bg-zinc-200',
        danger: 'bg-red-400 hover:bg-red-500 disabled:bg-red-200',
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
}
