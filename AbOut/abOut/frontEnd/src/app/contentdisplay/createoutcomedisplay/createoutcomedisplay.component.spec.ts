import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule} from '@angular/common/http/testing';
import { FormsModule } from '@angular/forms';

import { CreateoutcomedisplayComponent } from './createoutcomedisplay.component';

describe('CreateoutcomedisplayComponent', () => {
  let component: CreateoutcomedisplayComponent;
  let fixture: ComponentFixture<CreateoutcomedisplayComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CreateoutcomedisplayComponent ],
      imports: [RouterTestingModule, HttpClientTestingModule,FormsModule]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateoutcomedisplayComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
