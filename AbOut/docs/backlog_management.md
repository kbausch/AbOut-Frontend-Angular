# Backlog Management

This repository is a playground for setting up the scrum backlog environment 
in Gitlab.

## Primary Scrum Artifacts

Scrum divides artifiacts into three different class:

#### 1. Backlog

The backlog refers to the collection of user stories. The back log should be 
organized by priority by the product owner so that higher priority features 
may be tackled by the development team.

In Gitlab, the backlog is the list of issues labelled by priority by the 
product owner.

#### 2. Sprint

The sprint is associated with all tasks derived from user stores that belong 
to the sprint.

in Gitlab, the sprint is represented by project milestones.

#### 3. Issues

Issues are used to cover other items with different workflows. These are 
commonly bugs and defects. This class is not analygous to the issues construct 
on Gitlab.

Each of these constructs will be explored in depth below.

---

### User Stories

User stories are a high level overview of the functionality of the software 
from the perspective of the user. A user story has the general form:

```
As a <end-user-type>, I want to be able to <user-requirement> so that <reason>.
```

This structure can be misused to provide more information than needed. Bad 
user stories may be refined as the development team discusses them across the 
first few sprints. In this time, user stories of this nature should be resolved.

### Epics

User stories are meant to be independent of each other. There may be a case 
where some feature may depend on another. To deal with this, Epics are utilized. 
Epics can be thought of as a group of user stories.

Epics consist of a larger workflow where that specify an abstraction above user 
stories: Epics consist of a larger user flow consisting of multiple features.
In GitLab, an epic also contains a title and description, much like an issue, 
but it allows you to attach multiple child issues to it to indicate that 
hierarchy.

#### A Metric of Workload: User Story Points

User story points are assigned to a user story as an estimator for how much 
work a particular user story will require. It will require a few sprints to 
make this metric useful. These points are vital for estimating how many 
user stories can be taken on during a sprint. Consider the following scenario:

```
Team delivered:
        3 user stories in sprint 1 (total user story points 40)
        4 user stories in sprint 2 (total user story points 45)
        3 user stories in sprint 3 (total user story points 32)
```

The average of the user story points for each sprint is a measure known as 
velocity. The average of the three user story points per sprint above is 39. 
Given a user story worth 50 user story points, we may want to complete that 
story over the course of two sprints as it is much higher than the velocity of 
39.

In Gitlab, weights can be assigned for issues and closely mirrors the system of 
user story points.

#### Gitlab Equivalant

User stories correspond to issues on Gitlab. Each issue should identify a single 
user story.

---

### Tasks

A task is simply something to be completed by one person. It is differeniated 
from the user story by this fact: a user story is completed by many individuals, 
while a task is completed by a single individual. The user story is separated 
into many various tasks for completion by individuals.

#### Gitlab Equivalent

Since user stories are separated into various tasks, a task list can be 
created within an individual issue on Gitlab to further specify those 
individual tasks.

---

### Issues

Issues refer to anything that may have a different workflow than the user 
stories. Issues, like user stories, also belong in the project backlog. It is 
a common misunderstanding that only user stories may exist in the project 
backlog. The following quote is from the Scum wikipedia page:

```
There is a common misunderstanding that only user stories are allowed in a product backlog. By contrast, scrum is neutral on requirement techniques. As the Scrum Primer[19] states, Product Backlog items are articulated in any way that is clear and sustainable. Contrary to popular misunderstanding, the Product Backlog does not contain "user stories"; it simply contains items. Those items can be expressed as user stories, use cases, or any other requirements approach that the group finds useful. But whatever the approach, most items should focus on delivering value to customers.
```

#### Gitlab Equivalent

There is no set Gitlab equivalent for Scrum issues. As the previous section 
suggests, the representation of issues in Gitlab should be decided by the team 
so that the artifact is useful to the customers. A particular implementation 
for a bug issue could be a Gitlab issue explaning the bug and how to produce it 
identified with a bug label.

---

## References

[Scrum (software development) - Wikipedia](https://en.wikipedia.org/wiki/Scrum_(software_development)#Artifacts)

[User Stories Demystified](https://blog.taiga.io/user-stories-demystified.html)

[Difference Between a User Story and Task](https://www.mountaingoatsoftware.com/blog/the-difference-between-a-story-and-a-task)

[Issue vs User Story vs Task on Boards](https://github.com/taigaio/taiga-front/issues/797)

[How to use Gitlab for Agile Development](https://about.gitlab.com/blog/2018/03/05/gitlab-for-agile-software-development/)