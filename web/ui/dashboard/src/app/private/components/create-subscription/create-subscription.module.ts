import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateSubscriptionComponent } from './create-subscription.component';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAppModule } from '../create-app/create-app.module';
import { CreateSourceModule } from '../create-source/create-source.module';
import { TooltipModule } from '../tooltip/tooltip.module';
import { LoaderModule } from '../loader/loader.module';
import { CreateEndpointModule } from '../../pages/project/apps/app-details/create-endpoint/create-endpoint.module';
import { InputComponent } from 'src/app/components/input/input.component';
import { SelectComponent } from 'src/app/components/select/select.component';
import { ButtonComponent } from 'src/app/components/button/button.component';
import { TooltipComponent } from 'src/app/components/tooltip/tooltip.component';
import { ToggleComponent } from 'src/app/components/toggle/toggle.component';

@NgModule({
	declarations: [CreateSubscriptionComponent],
	imports: [
		CommonModule,
		ReactiveFormsModule,
		CreateAppModule,
		CreateSourceModule,
		TooltipModule,
		LoaderModule,
		CreateEndpointModule,
		InputComponent,
		SelectComponent,
		ButtonComponent,
		TooltipComponent,
		ToggleComponent
	],
	exports: [CreateSubscriptionComponent]
})
export class CreateSubscriptionModule {}
