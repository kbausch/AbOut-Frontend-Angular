import { TestBed, async, inject } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

import { AuthenticationService } from './authentication.service';

describe('AuthenticationService', () => {
  beforeEach(() => TestBed.configureTestingModule({
    imports: [
      HttpClientTestingModule
      ]
  }));

  it(`should create`, async(inject([HttpTestingController, AuthenticationService],
    (httpClient: HttpTestingController, authService: AuthenticationService) => {
      expect(authService).toBeTruthy();
  })));
});
