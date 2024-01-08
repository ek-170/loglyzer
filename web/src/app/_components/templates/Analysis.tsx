'use client';
import Button from '@/app/_components/atoms/buttons/Button';

export default function Analysis() {
  const buttonElementProps = {
    onClick: () => {
      console.log('button pushed!');
    },
    disabled: false,
  };

  return (
    <>
      <Button type="positive" buttonElementProps={buttonElementProps}>
        Save
      </Button>
    </>
  );
}
