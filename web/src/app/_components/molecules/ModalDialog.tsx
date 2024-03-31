/* eslint-disable tailwindcss/no-custom-classname */
import { Bottom, Top } from '@/app/_types/tailwind/TopBottomRightLeft';
import { ComponentPropsWithRef, ReactNode, forwardRef } from 'react';
import { tv } from 'tailwind-variants';

export type ModalDialogProps = {
  top?: Top;
  bottom?: Bottom;
  header?: ReactNode;
  children: ReactNode;
  isOpen: boolean;
} & Omit<ComponentPropsWithRef<'div'>, 'ref'>;

export const ModalDialog = forwardRef<HTMLDivElement, ModalDialogProps>(
  (props, ref) => {
    const { top, bottom, header, children, isOpen, ...divElementProps } = props;
    const modalTv = tv({
      base: `fixed ${top ?? 'top-0'} ${
        bottom ?? 'bottom-0'
      } left-0 flex h-dvh w-dvw items-center justify-center`,
      variants: {
        state: {
          open: 'animate-appear bg-black-600/75',
          close: 'invisible animate-disappear',
        },
      },
    });

    return (
      <>
        <div
          {...divElementProps}
          ref={ref}
          className={modalTv({ state: isOpen ? 'open' : 'close' })}
        >
          <div className="h-fit w-fit flex-col rounded-md bg-white">
            {header && <div>{header}</div>}
            <div className="overflow-auto p-3">{children}</div>
          </div>
        </div>
      </>
    );
  },
);

ModalDialog.displayName = 'ModalDialog';
