import { Story, Meta } from '@storybook/angular/types-6-0';
import { moduleMetadata } from '@storybook/angular';
import { InputComponent } from '../../app/components/input/input.component';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

export default {
	title: 'Example/Input',
	component: InputComponent,
	decorators: [
		moduleMetadata({
			imports: [ReactiveFormsModule]
		})
	],
	argTypes: {
		label: {
			control: { type: 'text' }
		},
		type: {
			options: ['text', 'password', 'number', 'url', 'email'],
			control: { type: 'select' },
			defaultValue: 'text'
		},
		errorMessage: {
			control: { type: 'text' }
		},
		placeholder: {
			control: { type: 'text' }
		},
		required: {
			control: { type: 'boolean' }
		},
		readonly: {
			control: { type: 'boolean' }
		},
		tooltipPosition: {
			options: ['left', 'right'],
			control: { type: 'select' },
			defaultValue: 'left'
		},
		tooltipSize: {
			options: ['sm', 'md'],
			control: { type: 'select' },
			defaultValue: 'md'
		},
		tooltipContent: {
			control: { type: 'text' }
		}
	},
	parameters: {
		actions: {
			handles: ['focus']
		}
	}
} as Meta;

const Template: Story<InputComponent> = (args: InputComponent) => {
	const formGroup = new FormGroup({
		name: new FormControl(undefined)
	});

	return {
		component: InputComponent,
		template: `
        <form [formGroup]="form" class="pt-100px">
          <convoy-input
            [label]="label"
            [type]="type"
            [placeholder]="placeholder"
            [errorMessage]="errorMessage"
            [required]="required"
            [readonly]="readonly"
            [tooltipPosition]="tooltipPosition"
            [tooltipSize]="tooltipSize"
            [tooltipContent]="tooltipContent"
            formControlName="name"
          >
          </convoy-input>
        </form>
      `,
		props: {
			...args,
			form: formGroup
		}
	};
};

export const Base = Template.bind({});
Base.args = {
	label: 'Convoy Input Label',
	type: 'text',
	placeholder: 'Convoy input placeholder',
	required: true,
	readonly: false,
	errorMessage: 'Convoy input error message',
	tooltipContent: 'Convoy input tooltip content'
} as Partial<InputComponent>;

export const Password = Template.bind({});
Password.args = {
	label: 'Convoy Input Label',
	type: 'password',
	placeholder: '********',
	required: true,
	readonly: false,
	errorMessage: 'Convoy input error message'
} as Partial<InputComponent>;
