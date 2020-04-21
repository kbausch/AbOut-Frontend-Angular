import { Component, OnInit } from '@angular/core';
import { HttpClient} from '@angular/common/http';
import { Outcome } from '../../models/outcome.model';
import { ToastService } from '../../services/toast/toast.service';

@Component({
  selector: 'app-outcomes-sidenavbar',
  templateUrl: './outcomes-sidenavbar.component.html',
  styleUrls: ['./outcomes-sidenavbar.component.scss']
})
export class OutcomesSidenavbarComponent implements OnInit {

  // The outcomeList and date will be used by the HTML to display the data.
  private outcomeList: Outcome[] = [];
  private currentDate = new Date();

  constructor(private http: HttpClient, private toastService: ToastService) { }

  ngOnInit() {
    this.getOutcomes();
  }

  // getOutcomes gets the JSON from the backend and stores it in outcomes_list.
  // It also displays an error message via the toast service if one occurs.
  private getOutcomes() {
    this.http.get<Outcome[]>('outcomes').subscribe(
      (response) => {
        this.outcomeList = response;
      },
      (error) => {
        this.toastService.show(error.status, error.message, { classname: 'bg-danger text-light' });
      }
    );
  }

}
