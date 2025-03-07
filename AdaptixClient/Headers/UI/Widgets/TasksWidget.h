#ifndef ADAPTIXCLIENT_TASKSWIDGET_H
#define ADAPTIXCLIENT_TASKSWIDGET_H

#include <main.h>
#include <Utils/CustomElements.h>

class TaskOutputWidget : public QWidget
{
    QGridLayout* mainGridLayout = nullptr;
    QLabel*      label          = nullptr;
    QLineEdit*   inputMessage   = nullptr;
    QTextEdit*   outputTextEdit = nullptr;

    void createUI();

public:
    explicit TaskOutputWidget( );
    ~TaskOutputWidget();

    void SetConten(QString message, QString text);
};



class TasksWidget : public QWidget
{
    QWidget*      mainWidget     = nullptr;
    QGridLayout*  mainGridLayout = nullptr;
    QTableWidget* tableWidget    = nullptr;
    QComboBox*    comboAgent     = nullptr;
    QComboBox*    comboStatus    = nullptr;
    QLineEdit*    inputFilter    = nullptr;

    void createUI();
    bool filterItem(TaskData task);
    void addTableItem(TaskData task);

public:
    TaskOutputWidget* taskOutputConsole = nullptr;

    explicit TasksWidget( QWidget* w );
    ~TasksWidget();

    void Clear();
    void AddTaskItem(TaskData newTask);
    void EditTaskItem(TaskData newTask);
    void RemoveTaskItem(QString taskId);
    void SetData();
    void SetAgentFilter(QString agentId);

public slots:
    void handleTasksMenu( const QPoint &pos );
    void onTableItemSelection();
    void onAgentChange(QString agentId);
    void actionCopyTaskId();
    void actionCopyCmd();
    void actionOpenConsole();
    void actionStop();
    void actionDelete();
};

#endif //ADAPTIXCLIENT_TASKSWIDGET_H
