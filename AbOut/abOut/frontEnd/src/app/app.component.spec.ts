import { TestBed, async } from '@angular/core/testing';
import { AppComponent } from './app.component';
import { ReactiveFormsModule } from '@angular/forms';
import { TopnavbarModule} from './topnavbar/topnavbar.module';
import { RouterTestingModule } from '@angular/router/testing';
import { OutcomesSidenavbarComponent } from './sidenavbar/outcomes-sidenavbar/outcomes-sidenavbar.component';
import { ProgramsSidenavbarComponent } from './sidenavbar/programs-sidenavbar/programs-sidenavbar.component';
import { GeneraldisplayComponent } from './contentdisplay/generaldisplay/generaldisplay.component';
import { ViewdisplayComponent } from './contentdisplay/viewdisplay/viewdisplay.component';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import { ToastsComponent } from './services/toast/toasts-container.component';
import { HttpClientTestingModule} from '@angular/common/http/testing';
import {APP_BASE_HREF} from '@angular/common';
import { ViewprogramdisplayComponent } from './contentdisplay/viewprogramdisplay/viewprogramdisplay.component';

describe('AppComponent', () => {
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        AppComponent,
        OutcomesSidenavbarComponent,
        ProgramsSidenavbarComponent,
        GeneraldisplayComponent,
        ViewdisplayComponent,
        ToastsComponent,
        ViewprogramdisplayComponent
      ],
      imports: [ReactiveFormsModule,
        TopnavbarModule,
        RouterTestingModule,
        NgbModule,
        HttpClientTestingModule],
      providers: [{ provide: APP_BASE_HREF, useValue : '/' }]
    }).compileComponents();
  }));

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have as title 'AbOut'`, () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('AbOut');
  });
});
