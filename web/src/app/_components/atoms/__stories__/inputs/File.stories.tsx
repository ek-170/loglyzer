import type { Meta, StoryObj } from '@storybook/react';

import { File } from '@/app/_components/atoms/inputs/File';

const meta: Meta<typeof File> = {
  component: File,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
};

export default meta;
type Story = StoryObj<typeof File>;

export const SampleFile: Story = {
  args: {
    label: 'file',
    color: 'light',
  },
};
