import type { Meta, StoryObj } from '@storybook/react';

import { PlusIcon } from '@/app/_components/atoms/icons/PlusIcon';

const meta: Meta<typeof PlusIcon> = {
  component: PlusIcon,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof PlusIcon>;

export const Icon: Story = {};
