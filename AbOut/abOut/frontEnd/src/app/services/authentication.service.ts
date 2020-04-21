import { Injectable, isDevMode } from '@angular/core';
import { HttpClient, HttpRequest } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {

  private authenticatedSubject: BehaviorSubject<boolean>;
  public authenticated: Observable<boolean>;
  private inProduction: boolean;

  constructor(private http: HttpClient) {
    const userData = localStorage.getItem('access_token');
    if (userData != null) {
      this.authenticatedSubject = new BehaviorSubject<boolean>(true);
    } else {
      this.authenticatedSubject = new BehaviorSubject<boolean>(false);
    }
    this.authenticated = this.authenticatedSubject.asObservable();
    this.inProduction = !isDevMode();
  }

  public isAuthenticated(): boolean {
    return this.authenticatedSubject.value;
  }

  // Attempts to log the user in from the fake login system.
  public fakeLogin(usernameCAS: string) {
    return this.http.get('auth/' + usernameCAS)
        .subscribe(
          (tokenData) => {
            const key = 'token';
            localStorage.setItem('access_token', tokenData[key]);
            this.authenticatedSubject.next(true);
          },
          (error) => {
            // Let it be.
          }
        );
  }

  // casLogin will attempt to authenticate the user from CAS.
  public casLogin(usernameCAS: string, password: string) {
    return this.http.get('auth/login')
        .subscribe(
          (res) => {
            // If there is no error, we are authenticated.
            this.authenticatedSubject.next(true);
          },
          (error) => {
            // We are not authenticated, so the response will 302 and try
            // to redirect us.
          }
        );
  }

  // Clears token information from the local storage.
  public logout() {
    // Test for fake login in development mode.
    if (!this.inProduction) {
      localStorage.clear();
      this.authenticatedSubject.next(false);
    } else {
      // CAS logout for production.
      return this.http.get('auth/logout')
        .subscribe(
          (res) => {
            // No redirect, so no logout occured.
          },
          (error) => {
            // Error should be 302, we are redirected to the logout page and 
            // will clear the local storage and set the subjects.
            localStorage.clear();
            this.authenticatedSubject.next(false);
          }
        );
    }
  }
}
