import { Height, Width } from '@/app/_types/tailwind/Sizing';
import { Padding, Margin } from '@/app/_types/tailwind/Spacing';
import { ComponentPropsWithRef, ReactNode } from 'react';
import { tv } from 'tailwind-variants';

/**
 * padding and width should be written in the way the class name of tailwind is written
 */
export type ButtonProps = {
  color?: 'positive' | 'negative' | 'danger' | 'primary';
  padding?: Padding[];
  margin?: Margin[];
  height?: Height;
  width?: Width;
  children?: ReactNode;
} & ComponentPropsWithRef<'button'>;

export const Button = (props: ButtonProps) => {
  const {
    color,
    padding,
    margin,
    height,
    width,
    children,
    ...buttonElementProps
  } = props;
  const paddings = padding && padding.length > 0 ? padding : ['px-3', 'py-1.5'];
  const margins = margin && margin.length > 0 ? margin : [];
  const paddingClass = paddings.join(' ');
  const marginClass = margins.join(' ');
  const heightClass = height || 'h-fit';
  const widthClass = width || 'w-fit';
  const button = tv({
    base: `flex rounded-md ${paddingClass} ${marginClass} ${heightClass} ${widthClass} focus-visible:outline`,
    variants: {
      color: {
        positive: 'bg-sky-500 shadow-sm hover:bg-sky-300 disabled:bg-sky-100',
        negative:
          'bg-black-400 shadow-sm hover:bg-black-500 disabled:bg-black-200',
        danger: 'bg-red-500 shadow-sm hover:bg-red-400 disabled:bg-red-200',
        primary:
          'bg-primary-500 shadow-sm hover:bg-primary-400 disabled:bg-primary-200',
      },
    },
  });
  return (
    <>
      <button {...buttonElementProps} className={button({ color: color })}>
        {children}
      </button>
    </>
  );
};
