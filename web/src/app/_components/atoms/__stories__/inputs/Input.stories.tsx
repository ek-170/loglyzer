import type { Meta, StoryObj } from '@storybook/react';
import { useState } from 'react';
import { Input } from '@/app/_components/atoms/inputs/Input';

const meta: Meta<typeof Input> = {
  component: Input,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;

type Story = StoryObj<typeof Input>;

export const TextInput: Story = {
  args: {
    type: 'text',
    label: 'user name',
    color: 'light',
  },
};

export const DisabledInput: Story = {
  args: {
    type: 'text',
    label: 'user name',
    color: 'light',
    disabled: true,
  },
};

export const InvalidInput: Story = {
  args: {
    label: 'Email',
    color: 'light',
  },
  render: function Wrapper(...args) {
    const [value, setValue] = useState('test@gmail.com.');
    return (
      <Input
        {...args[0]}
        type={'email'}
        color={'light'}
        value={value}
        onChange={(e) => setValue(e.target.value)}
      />
    );
  },
};
