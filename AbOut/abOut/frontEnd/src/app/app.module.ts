import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { TopnavbarModule} from './topnavbar/topnavbar.module';

import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { JwtModule } from '@auth0/angular-jwt';
import { environment } from 'src/environments/environment';
import { BaseurlService } from './services/baseurl.service';

import { AppRoutingModule } from './app-routing.module';

import { OutcomesSidenavbarComponent } from './sidenavbar/outcomes-sidenavbar/outcomes-sidenavbar.component';
import { ProgramsSidenavbarComponent } from './sidenavbar/programs-sidenavbar/programs-sidenavbar.component';

import { GeneraldisplayComponent } from './contentdisplay/generaldisplay/generaldisplay.component';
import { ViewdisplayComponent } from './contentdisplay/viewdisplay/viewdisplay.component';
import { CreateoutcomedisplayComponent } from './contentdisplay/createoutcomedisplay/createoutcomedisplay.component';

import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { ToastsComponent } from './services/toast/toasts-container.component';

import {APP_BASE_HREF} from '@angular/common';
import { ViewprogramdisplayComponent } from './contentdisplay/viewprogramdisplay/viewprogramdisplay.component';


// Fetches authentication token from the local storage.
export function tokenGetter() {
  return localStorage.getItem('access_token');
}

@NgModule({
  declarations: [
    AppComponent,
    ViewdisplayComponent,
    OutcomesSidenavbarComponent,
    ProgramsSidenavbarComponent,
    GeneraldisplayComponent,
    ToastsComponent,
    CreateoutcomedisplayComponent,
    ViewprogramdisplayComponent
  ],
  imports: [
    TopnavbarModule,
    NgbModule,
    BrowserModule,
    HttpClientModule,
    FormsModule,
    AppRoutingModule,
    JwtModule.forRoot({
      config: {
        tokenGetter,
        whitelistedDomains: ['localhost:4200', 'localhost:8080', 'mtlbsso.mtech.edu']
      }
    }),
    ReactiveFormsModule
  ],
  providers: [
    // BaseurlService injects the base url into all http requests.
    { provide: HTTP_INTERCEPTORS, useClass: BaseurlService, multi: true },
    // This provider sets the value for the base url.
    { provide: 'BASE_API_URL', useValue: environment.ApiUrl },
    // This sets the base href
    { provide: APP_BASE_HREF, useValue : '/' }
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
