/* eslint-disable tailwindcss/no-custom-classname */
import { Bottom, Top } from '@/app/_types/tailwind/TopBottomRightLeft';
import { ComponentPropsWithRef, ReactNode } from 'react';
import { tv } from 'tailwind-variants';

export type ModalProps = {
  top?: Top;
  bottom?: Bottom;
  children: ReactNode;
  isOpen: boolean;
} & ComponentPropsWithRef<'div'>;;

export const Modal = (props: ModalProps) => {
  const {
    top,
    bottom,
    children,
    isOpen,
    ...divElementProps
  } = props;
  const modalTv = tv({
    base: `relative ${top ?? "top-0"} ${bottom ?? "bottom-0"} left-0 h-dvh w-dvw flex items-center justify-center`,
    variants: {
      state: {
        open: 'bg-black-500 animate-appear opacity-60	',
        close: 'animate-disappear invisible',
      },
    },
  });
  return (
    <>
      <div {...divElementProps} className={modalTv({state: isOpen ? "open" : "close"})}>
        {children}
      </div>
    </>
  );
};
