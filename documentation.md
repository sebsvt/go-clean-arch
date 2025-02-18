# Clean Architecture in Golang

Clean architecture เป็นแนวคิดในการเขียนโค้ดอย่างนึงที่นำเสนอแนวคิดที่ว่า Business Domain มันควรที่จะตัดขาดหรือไม่ควรมี dependency กับ framework หรือ tools ใดๆ

หากเราดู tutorial ใน yt ต่างๆ architecture ที่เราจะได้เห็นบ่อยๆก็คือ MVC
ทุกๆ logic และ โค้ดในการติดต่อ database ต่างๆจะกองกันอยู่ที่ function ของ endpoint นั้นๆหรือหรูๆหน่อยจะมีการแยก function usecase เอาไว้แยกต่างหากเพื่อให้มัน reuseable ได้แต่ใน function มันก็มี depedency กับ external tools อยู่ดีหากเราเปลี่ยน database หรือ api framework กลายเป็นว่าแทนที่จะเปลี่ยนแค่โค้ดของ database / api framework กับกลายเป็นการเขียนใหม่ทั้งหมดเพราะโค้ดตรงนี้มันมีการเรียกจากตรงนู้นโค้ดตรงนู้นมีการเรียกจากตรงนี้จนกลายเป็น Big Ball of Mud และยากต่อการแยกโค้ดส่วนต่างๆมาทำการเทส

ดังนั้น Uncle Bob ของเราเลยเสนอว่า เออจากปัญหาตรงนี้เนี่ยทำไมเราไม่แยก business logic ออกจาก technology/framework ล่ะโดยจะมีการแบ่ง layers ออกเป็นตามนี้

- Models Layer
- Repository Layer
- Usecase Layer
- Delivery Layer

Models Layer เนี่ยคือทีๆเราไว้นิยาม Business Domain Model ของเราเช่น User ควรมีหน้าตาอย่างไรหรือ Account ควรมีข้อมูลอะไรบ้าง

Repository Layer คือที่เราไว้สร้างหรือเขียนโค้ดใช้งาน framework ภายนอกโดยเราจะมีการนิยาม Interface เอาไว้ เวลาพัฒนาจะได้เป็น Protocal เดียวกันว่าควรมี Methods หรือ functions อะไรเพราะถ้าไม่มีการนิยามเอาไว้ต่างฝ่ายก็ต่างนิยามชื่อฟังชั่นมั่วซั่วลองจินตาการว่าทุกวันนี้เราไม่มีภาษาอังกฤษเป็นภาษากลางและทุกคนเอาแต่พูดภาษาตัวเองดูสิ คงปวดหัวแน่นอน โดยส่วนตัวแล้วผมเรียกว่า Adapter มากกว่า Repository ฮ่าๆโดยที่ Interface คือ Port

Usecase Layer ณ ที่แห่งนี้มันคือแหล่งที่เราไว้เขียน Business Logic โดยใช้ Model ของ Business Domain Model และที่ตรงนี้นี่เองเราจะมีการเรียกใช้งาน external tools ผ่าน interface เช่น Account ใน Business model อาจจะมีฟิลด์เกี่ยวกับ เงินในบัญชี, ชื่อ, id โดยที่ usecase จะประกอบไปด้วย logic หรือ usecase ที่เกี่ยวกับบัญชีเช่น deposit, withdraw

Delivery Layer คือ layer ที่ทำหน้าที่ติดต่อกับผู้ใช้ไม่ว่าจะเป็นผ่าน API, CLI หรือ gRPC เท่านั้น

ในโปรเจคนี้จะมีการแบ่งตาม Clean Architecture โดยที่ usecase layer ของเราจะเป็น package ใหม่เลยที่เรานิยามไว้เช่น user, account ... ในขณะที่ domain คือ Model layer ของเรา ส่วน internal จะเป็นส่วนที่เราเก็บพวก repository, rest(delivery layer)
