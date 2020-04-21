import { Component, isDevMode } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { AuthenticationService } from './services/authentication.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  private title = 'AbOut';
  public loginForm: FormGroup;
  public authenticated = false;
  private authSubscription: Subscription;
  public inProduction: boolean;

  constructor(
    private formBuilder: FormBuilder,
    private authService: AuthenticationService,
  ) {
    this.loginForm = this.formBuilder.group({
      usernameCAS: '',
      password: '',
    });
    this.authenticated = this.authService.isAuthenticated();
    this.authSubscription = authService.authenticated.subscribe(
      (val) => {
        this.authenticated = val;
      }
    );
    this.inProduction = !isDevMode();
  }

  public onFakeLogin(loginData) {
    // The fake login will only be active if dev mode is active.
    if (this.loginForm.invalid) {
      return;
    }

    const username = loginData.usernameCAS;
    // This request will always yield a valid token, as long as the username
    // is not empty.
    this.authService.fakeLogin(username);
  }

  public onCasLogin(loginData) {
    if (this.loginForm.invalid) {
      return;
    }

    const username = loginData.usernameCAS;
    const password = loginData.password;

    this.authService.casLogin(username, password)
  }

  public onLogout() {
    this.authService.logout();
  }

}
