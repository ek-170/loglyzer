import type { Meta, StoryObj } from '@storybook/react';

import { CrossIcon } from '@/app/_components/atoms/icons/CrossIcon';

const meta: Meta<typeof CrossIcon> = {
  component: CrossIcon,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof CrossIcon>;

export const Icon: Story = {};
