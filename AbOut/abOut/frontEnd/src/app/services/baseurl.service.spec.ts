import { TestBed } from '@angular/core/testing';
import { environment } from 'src/environments/environment';

import { BaseurlService } from './baseurl.service';

describe('BaseurlService', () => {
  beforeEach(() => TestBed.configureTestingModule({
    providers: [{ provide: 'BASE_API_URL', useValue: environment.ApiUrl }]
  }));

  it('should be created', () => {
    const service: BaseurlService = TestBed.get(BaseurlService);
    expect(service).toBeTruthy();
  });
});
