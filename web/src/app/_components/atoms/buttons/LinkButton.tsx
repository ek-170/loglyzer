import { Width } from '@/app/_types/tailwind/Sizing';
import { Padding, Margin } from '@/app/_types/tailwind/Spacing';
import { ReactNode } from 'react';
import Link, { LinkProps } from 'next/link';
import { tv } from 'tailwind-variants';

/**
 * padding and width should be written in the way the class name of tailwind is written
 */
export type LinkButtonProps = {
  color: 'primary' | 'elastic';
  padding?: Padding[];
  margin?: Margin[];
  width?: Width;
  children?: ReactNode;
} & LinkProps;

export const LinkButton = (props: LinkButtonProps) => {
  const { color, padding, margin, width, children, ...linkElementProps } =
    props;
  const paddings = padding && padding.length > 0 ? padding : ['px-3', 'py-1.5'];
  const margins = margin && margin.length > 0 ? margin : [];
  const paddingClass = paddings.join(' ');
  const marginClass = margins.join(' ');
  const widthClass = width || 'w-fit';
  const wrapper = tv({
    base: `rounded-md ${widthClass}`,
    variants: {
      color: {
        elastic: 'bg-elastic-400 hover:bg-elastic-200 disabled:bg-elastic-100',
        primary: 'bg-primary-500 hover:bg-primary-400 disabled:bg-primary-200',
      },
    },
  });
  const link = tv({
    base: `flex justify-center rounded-md ${paddingClass} ${marginClass} w-full`,
  });
  return (
    <div className={wrapper({ color: 'elastic' })}>
      <Link {...linkElementProps} className={link()}>
        {children}
      </Link>
    </div>
  );
};
