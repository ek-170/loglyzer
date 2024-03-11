/* eslint-disable tailwindcss/no-custom-classname */
import { Bottom, Top } from '@/app/_types/tailwind/TopBottomRightLeft';
import { ComponentPropsWithRef, ReactNode } from 'react';
import { tv } from 'tailwind-variants';

export type ModalProps = {
  top?: Top;
  bottom?: Bottom;
  children: ReactNode;
  isOpen: boolean;
} & ComponentPropsWithRef<'div'>;

export const Modal = (props: ModalProps) => {
  const { top, bottom, children, isOpen, ...divElementProps } = props;
  const modalTv = tv({
    base: `relative ${top ?? 'top-0'} ${
      bottom ?? 'bottom-0'
    } left-0 flex h-dvh w-dvw items-center justify-center`,
    variants: {
      state: {
        open: 'animate-appear bg-black-500 opacity-60	',
        close: 'invisible animate-disappear',
      },
    },
  });
  return (
    <>
      <div
        {...divElementProps}
        className={modalTv({ state: isOpen ? 'open' : 'close' })}
      >
        {children}
      </div>
    </>
  );
};
