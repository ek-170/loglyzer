import { Button, ButtonProps } from '../atoms';
import { ReactElement } from 'react';
import { SvgIconProps } from '../atoms/icons/SvgIcon';

type IconButtonProps<T extends React.ComponentType<SvgIconProps>> =
  ButtonProps & {
    startIcon?: ReactElement<React.ComponentProps<T>>;
    endIcon?: ReactElement<React.ComponentProps<T>>;
  };

export const IconButton = (
  props: IconButtonProps<React.ComponentType<SvgIconProps>>,
) => {
  const {
    startIcon: startIcon,
    endIcon: endIcon,
    children: children,
    ...buttonProps
  } = props;
  return (
    <>
      <Button {...buttonProps}>
        {startIcon && <div className="p-px pr-1.5">{startIcon}</div>}
        {children && <div className="px-px">{children}</div>}
        {endIcon && <div className="p-px pl-1.5">{endIcon}</div>}
      </Button>
    </>
  );
};
