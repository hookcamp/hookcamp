import { ENDPOINT } from './endpoint.model';
import { SOURCE } from './group.model';

export interface SUBSCRIPTION {
	created_at: string;
	endpoint: string;
	name: string;
	source: SOURCE;
	status: string;
	type: 'outgoing' | 'incoming';
	uid: string;
	updated_at: string;
	endpoint_metadata?: ENDPOINT;
	alert_config?: { count: number; threshold: string };
	retry_config?: { type: string; retry_count: number; duration: number };
	source_metadata: SOURCE;
	filter_config: { event_types: string[]; filter: { headers: string; body: string } };
	active_menu?: boolean;
	device_metadata: { created_at: string; deleted_at: string; host_name: string; last_seen_at: string; status: 'offline' | 'online'; uid: string; updated_at: string };
}
