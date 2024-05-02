# CS-506-Predictive-Kubernetes

Project Summary:
This project focuses on developing a predictive resource scaling system for Kubernetes clusters. As organizations increasingly rely on Kubernetes for infrastructure management, optimizing resource utilization becomes crucial for cost-effectiveness and performance. Our project aims to address this need by creating a system that automatically adjusts resource allocation based on anticipated demand, thus improving efficiency and scalability.

Data Collection: Gathered CPU and memory usage data from local Kubernetes clusters running a sample application. Utilized various tools and scripts for real-time metric collection. Our entire data collection workflow including instructions on how to reproduce our sample data is in the sp24-team-a/weekly/week2 directory.

Data Analysis: Conducted extensive analysis to understand application resource usage patterns and develop predictive models. Explored various models and techniques, each of which are documented along with the relevant data sets in the week3 - week8 directories in sp24-team-a/weekly.

Kubernetes Integration: Integrated predictive models into Kubernetes using a hybrid approach of both a Python script and Kubernetes operator. Created custom resources to communicate scaling requirements to Kubernetes operators. This system can be found with the relevant instructions in the sp24-team-a/deliverables/predkube_operator directory.

Team Information:
Abir Islam: aislam99@bu.edu
Bowen Li: bown@bu.edu
Hyunjin Jung: hyunjin@bu.edu
Mark Yang: mdyang@bu.edu
William Seong: seongw@bu.edu