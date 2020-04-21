import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule} from '@angular/common/http/testing';
import { ReactiveFormsModule } from '@angular/forms';

import { ViewdisplayComponent } from './viewdisplay.component';

describe('ViewdisplayComponent', () => {
  let component: ViewdisplayComponent;
  let fixture: ComponentFixture<ViewdisplayComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewdisplayComponent ],
      imports: [RouterTestingModule, HttpClientTestingModule, ReactiveFormsModule]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewdisplayComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
